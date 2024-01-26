# DirWizard CLI

`DirWizard` is a command-line interface (CLI) tool developed by Kevin Liu, designed to simplify file system operations such as renaming directories, checking directory naming compliance, finding duplicate directories, and more. It provides an easy-to-use interface for interacting with traditional Linux CLI tools and shell scripts.

## Features

- **Rename Directories:** Convert all upper-case characters to lower-case, replace spaces and dashes with underscores, and clean up leading/trailing characters.
- **Check Compliance:** Verify directory naming conventions and log any non-compliant directory names.
- **Find Duplicates:** Identify and log duplicate directories based on file and directory names.
- **Generate Mock File Structure:** For testing purposes, create a mock file structure with various directory and file names.
- **Clear Mock File Structure:** Remove the generated mock file structure.

## Getting Started

### Prerequisites

- A Linux or Unix-like operating system.
- Go version 1.xx or higher.
- Basic knowledge of command-line operations.

### Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/kevinliu-bmb/DirWizard.git
   cd DirWizard
   ```

2. **Build the Project:**

   ```bash
   go build -o DirWizard
   ```

3. **Run the Program:**

   ```bash
   cd ./DirWizard
   ./DirWizard
   ```

### Usage

After running the program, you will be presented with a menu of operations. Enter the number corresponding to the operation you wish to perform and follow the on-screen instructions.

Example:

   ```bash
   Enter your choice (or 0 to exit): 1
   ```

This will initiate the directory renaming process.

## Contributing

Contributions to DirWizard are welcome! Feel free to fork the repository and submit pull requests.

## License

This project is licensed under the GNU General Public License v2.0 - see the [LICENSE](LICENSE) file for details.
