# CanUC Project

A multi-component project with C core and integrated submodules for learning management system functionality.

## Project Structure

```
CanUC/
├── src/                  # C source code
│   ├── main.c            # Main application entry point
│   ├── service/          # Service layer for business logic
│   ├── utils/            # Utilities submodule
│   └── app/              # Application interface layer
├── leb2/                 # leb2 submodule
├── .gitignore            # Git ignore file
└── .gitmodule            # Git submodules configuration
```

## Overview

CanUC is a modular project designed for managing course materials, calendars, and academic resources in a learning environment. The project leverages a C-based core with several specialized submodules to handle various functionalities including authentication, calendar management, and learning environment features.

## Components

### src (Core C Implementation)
- **main.c**: Entry point for the application
- **service/**: Handles business logic including authentication, data loading, and database management
- **utils/**: Common utilities, data types, and core data structures
- **app/**: User interface layer providing interactive pages for different features

### leb2 (Submodule)
External submodule that provides learning environment and backpack functionality for course management, assignments, and academic resources.

## Prerequisites

### C Compiler Requirement
This project requires **GCC version 14.1.0 or newer**. 

✅ **How to check your GCC version:**
Run this command in your terminal to see what version you have:
```bash
gcc --version
```

❌ **If your version is below 14.1.0:**
You'll need to update your compiler before building this project. See the "Compiler Installation" section below.

⚠️ **Why the version matters:**
The project uses C language features only available in recent GCC versions. Using older versions will result in compilation errors.

### Compiler Installation

**If you don't have GCC installed or need to update:**

- **Windows Users:** 
  Download the latest MinGW-w64 installer from their [official website](https://www.mingw-w64.org/downloads/).
  
- **Mac Users:** 
  Install using Homebrew: `brew install gcc`
  
- **Linux Users:** 
  Ubuntu/Debian: `sudo apt update && sudo apt install gcc`
  Fedora: `sudo dnf install gcc`

## Setting Up the Project

### 1. Clone the Repository with Submodules

```bash
git clone --recursive https://github.com/kmood-Sakura/CanUC.git
cd CanUC
```

### 2. Initialize Submodules
If you've already cloned without the `--recursive` flag, initialize the submodules:

```bash
git submodule update --init --recursive
```

### 3. Update Submodules
To ensure you have the latest version of all submodules:

```bash
git submodule update --remote --recursive
```

## Compilation and Running

### Build Process
```bash
go run ./compile.go
```

### Running the Application
After compilation, run the program:

**Linux/macOS:**
```bash
./build/program
```

**Windows:**
```bash
build\program.exe
```

## Usage

When you start the application, you'll see the authentication screen where you can create an account or log in to an existing one.

The main features include:
- **LEB2**: Manage your courses, assignments, and learning materials
- **Calendar**: Schedule and manage your tasks and appointments
- **Notifications**: View important updates and reminders

For detailed usage instructions, please refer to the documentation in src/README.md.

## Contributing

When contributing to this project:
1. Follow the established code structure and style
2. Use proper memory management (allocation, initialization, freeing)
3. Document new functions and structures with clear comments
4. Include comprehensive error handling for all operations
5. Test thoroughly before submitting changes

## License

[Specify your license information here]