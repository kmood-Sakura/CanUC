package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

type CompileOptions struct {
	SourceDir     string
	OutputDir     string
	IncludeDirs   []string
	CompilerFlags string
	OutputBinary  string
	Verbose       bool
	Jobs          int
}

func main() {
	// Define command line flags
	options := CompileOptions{}
	
	flag.StringVar(&options.SourceDir, "src", "src", "Root source directory")
	flag.StringVar(&options.OutputDir, "out", "build", "Output directory for object files")
	flag.StringVar(&options.CompilerFlags, "flags", "-Wall -Werror", "Compiler flags for gcc")
	flag.StringVar(&options.OutputBinary, "bin", "program", "Name of the output binary")
	includeFlag := flag.String("include", "", "Include directories (comma separated)")
	flag.BoolVar(&options.Verbose, "verbose", false, "Enable verbose output")
	flag.IntVar(&options.Jobs, "j", 4, "Number of parallel compilation jobs")
	flag.Parse()

	// Process include directories
	if *includeFlag != "" {
		options.IncludeDirs = strings.Split(*includeFlag, ",")
	}

	// Make sure output directory exists
	if err := os.MkdirAll(options.OutputDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Find all C files
	cFiles, err := findCFiles(options.SourceDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding C files: %v\n", err)
		os.Exit(1)
	}

	if len(cFiles) == 0 {
		fmt.Println("No C files found in the source directory.")
		os.Exit(0)
	}

	if options.Verbose {
		fmt.Printf("Found %d C files to compile\n", len(cFiles))
	}

	// Compile C files
	objFiles, err := compileFiles(cFiles, options)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Compilation failed: %v\n", err)
		os.Exit(1)
	}

	// Link object files
	if err := linkFiles(objFiles, options); err != nil {
		fmt.Fprintf(os.Stderr, "Linking failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully compiled and linked to %s\n", filepath.Join(options.OutputDir, options.OutputBinary))
}

func findCFiles(sourceDir string) ([]string, error) {
	var cFiles []string

	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), ".c") {
			cFiles = append(cFiles, path)
		}
		return nil
	})

	return cFiles, err
}

func compileFiles(cFiles []string, options CompileOptions) ([]string, error) {
	var objFiles []string
	var mutex sync.Mutex
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, options.Jobs)
	errors := make(chan error, len(cFiles))

	for _, cFile := range cFiles {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			semaphore <- struct{}{} // Acquire semaphore
			defer func() { <-semaphore }() // Release semaphore

			// Determine output object file path
			relPath, err := filepath.Rel(options.SourceDir, file)
			if err != nil {
				errors <- fmt.Errorf("failed to get relative path for %s: %v", file, err)
				return
			}

			objDir := filepath.Join(options.OutputDir, filepath.Dir(relPath))
			if err := os.MkdirAll(objDir, 0755); err != nil {
				errors <- fmt.Errorf("failed to create directory %s: %v", objDir, err)
				return
			}

			objFile := filepath.Join(objDir, filepath.Base(file[:len(file)-2]+".o"))

			// Generate include flags
			var includeFlags []string
			for _, dir := range options.IncludeDirs {
				includeFlags = append(includeFlags, "-I"+dir)
			}

			// Add include flag for the source directory to help with relative includes
			includeFlags = append(includeFlags, "-I"+options.SourceDir)

			// Add include flag for the file's own directory
			includeFlags = append(includeFlags, "-I"+filepath.Dir(file))

			// Compile the file
			args := []string{"-c", file, "-o", objFile}
			args = append(args, strings.Fields(options.CompilerFlags)...)
			args = append(args, includeFlags...)

			cmd := exec.Command("gcc", args...)
			
			if options.Verbose {
				fmt.Printf("Compiling %s -> %s\n", file, objFile)
				fmt.Println("Command:", cmd.String())
			}

			output, err := cmd.CombinedOutput()
			if err != nil {
				errors <- fmt.Errorf("compilation of %s failed: %v\n%s", file, err, string(output))
				return
			}

			mutex.Lock()
			objFiles = append(objFiles, objFile)
			mutex.Unlock()
		}(cFile)
	}

	wg.Wait()
	close(errors)

	// Check if any errors occurred
	for err := range errors {
		return nil, err
	}

	return objFiles, nil
}

func linkFiles(objFiles []string, options CompileOptions) error {
	outputPath := filepath.Join(options.OutputDir, options.OutputBinary)
	
	args := append([]string{"-o", outputPath}, objFiles...)
	
	// Add any linking flags if needed
	args = append(args, strings.Fields(options.CompilerFlags)...)
	
	cmd := exec.Command("gcc", args...)
	
	if options.Verbose {
		fmt.Println("Linking object files to", outputPath)
		fmt.Println("Command:", cmd.String())
	}
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("linking failed: %v\n%s", err, string(output))
	}
	
	return nil
}