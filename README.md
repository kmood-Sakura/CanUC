# CanUC Project

A multi-component project with C core and integrated submodules for learning management system functionality.

## Project Structure

```
CanUC/
├── src/                  # C source code
│   ├── main.c            # Main application entry point
│   ├── service/          # Service layer for business logic
│   │   ├── authenticate.h/c       # Authentication services
│   │   ├── calendar-service.h/c   # Calendar management
│   │   ├── load-user-data.h/c     # Data loading utilities
│   │   └── update-db.h/c          # Database update services
│   ├── utils/            # Utilities submodule
│   │   ├── common/       # Common utilities
│   │   │   ├── log.h/c   # Logging functionality
│   │   │   ├── request.h/c # User input handling
│   │   │   └── status.h/c # Status and error handling
│   │   ├── datatype/     # Data type definitions
│   │   │   ├── char-type.h/c  # Character utilities
│   │   │   ├── date-type.h/c  # Date/time handling
│   │   │   ├── float-type.h/c # Floating point types
│   │   │   ├── int-type.h/c   # Integer type definitions
│   │   │   └── string-type.h/c # String handling
│   │   └── struct/       # Core data structures
│   │       ├── auth.h/c  # Authentication structure
│   │       ├── calendar.h/c # Calendar structure
│   │       ├── file.h/c  # File operations
│   │       ├── leb2.h/c  # Learning environment
│   │       ├── notification.h/c # Notification system
│   │       └── path.h/c  # Path management
│   └── app/              # Application interface layer
│       ├── calendar-page.h/c   # Calendar UI
│       ├── home-page.h/c       # Main application UI
│       ├── leb2-page.h/c       # Learning environment UI
│       ├── login-page.h/c      # Authentication UI
│       └── notification-page.h/c # Notification UI
├── leb2/                 # leb2 submodule
├── .gitignore            # Git ignore file
├── .gitmodule            # Git submodules configuration
└── compile.go            # Go compilation script
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

## Getting Started

### Prerequisites
- C compiler (gcc/clang recommended)
- Go compiler (for build process)
- Git (for submodule management)

##### Setup Section

#### 1. Check GCC Installation
This project requires **GCC version 14.1.0 or newer**. 

**On Linux:**
```bash
gcc --version
```

**On macOS:**
```bash
gcc --version
```

**On Windows:**
```cmd
gcc --version
```

❌ **If your version is below 14.1.0:**
You'll need to update your compiler before building this project. See the "Compiler Installation" section below.

⚠️ **Why the version matters:**
The project uses C language features only available in recent GCC versions. Using older versions will result in compilation errors.

**On Ubuntu/Debian:**
```bash
sudo apt update
sudo apt install build-essential
```

**On macOS (using Homebrew):**
```bash
brew install gcc
```

**On Windows:**
- Download and install MinGW from [MinGW-w64](https://www.mingw-w64.org/downloads/)
- Add MinGW bin directory to your PATH environment variable

#### 2. Check Go Installation

```bash
go version
```

If Go is not installed:

**On Ubuntu/Debian:**
```bash
sudo apt update
sudo apt install golang-go
```

**On macOS (using Homebrew):**
```bash
brew install go
```

**On Windows:**
- Download the installer from [golang.org](https://golang.org/dl/)
- Run the installer and follow the instructions

### Installation and Submodule Setup

#### 1. Initialize Project

There are two ways to initialize the project:

**Option A: Clone the complete repository:**
```bash
git clone https://github.com/kmood-Sakura/CanUC.git
cd CanUC
```

**Option B: Create a new repository with the source code:**
```bash
# Create a new folder for your project
mkdir CanUC-Project
cd CanUC-Project

# Initialize git repository
git init

# Add the remote repository
git remote add origin https://github.com/kmood-Sakura/CanUC-src
git fetch
git checkout main
```

#### 2. Initialize Submodules

After cloning or initializing the repository, set up the submodules:

```bash
# Inside the CanUC directory
git submodule update --init --recursive
```

This command:
- Initializes your local configuration file for each submodule
- Fetches all the data from the submodule projects
- Checks out the appropriate commit specified in the superproject

#### 3. Update Submodules

To update submodules to their latest versions:

```bash
# Inside the CanUC directory
git submodule update --remote --recursive
```

This command:
- Fetches and updates all submodules to their latest versions from their respective repositories
- Updates nested submodules recursively
- Keeps your submodules in sync with upstream changes

##### Compilation Section

### Using Go Build Script
```bash
go run ./compile.go
```

This script will:
- Check for dependencies
- Compile all C source files
- Link the executable
- Place the output in the `build` directory

### Running the Application

**On Linux/macOS:**
```bash
./build/program
```

**On Windows:**
```cmd
build\program.exe
```

## Usage Examples

### Authentication Flow

```
--------------------------------------------------------

Authentication

  [1] Login
  [2] Sign up

  [e] Exit

  If you already have an account please login... please 🥺

command: 2

--------------------------------------------------------

  Sign up page

  Student ID: 12345678
  Password: password123
```

After successful sign-up, the system creates the user's directory structure and enters the main menu.

### Main Navigation

```
--------------------------------------------------------

Home

  [1] LEB2
  [2] Calendar
  [3] Notification

  [e] Exit

command: 
```

For more detailed usage examples, please see the documentation within the src/README.md file.