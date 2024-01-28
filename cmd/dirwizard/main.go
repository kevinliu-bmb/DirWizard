// This file is part of DirWizard.
// Copyright (C) 2024 Kevin Liu

// DirWizard is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 2 of the License, or
// (at your option) any later version.

// DirWizard is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with DirWizard.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var workingDir string

func main() {
    clearScreen()
    var err error
    workingDir, err = filepath.Abs(".")
    if err != nil {
        fmt.Println("Error getting current directory:", err)
        return
    }
    displayHeader()
    processUserInput()
}

func clearScreen() {
    if runtime.GOOS == "windows" {
        executeCommand("cmd", "/c", "cls")
    } else {
        executeCommand("clear")
    }
}

func executeCommand(name string, arg ...string) {
    cmd := exec.Command(name, arg...)
    cmd.Stdout = os.Stdout
    err := cmd.Run()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error executing command %s: %v\n", name, err)
    }
}

func displayHeader() {
    header := fmt.Sprintf(`=================================================
|         DirWizard - Directory Management      |
|         Version 0.1.0                         |
|         © 2024 Kevin Liu. All Rights Reserved |
=================================================
Working Directory: %s

`, filepath.Clean(workingDir))
    fmt.Print(header)
}

func displayMenu() {
    menu := `Options:

(1) - Rename Directories
(2) - Check Compliance
(3) - Find Duplicates
(4) - Search Data Paths

(5) - [dev test] Generate Mock File Structure
(6) - [dev test] Clear Mock File Structure

(7) - Change Working Directory

(0) - Exit

`
    fmt.Print(menu)
}

func processUserInput() {
    reader := bufio.NewReader(os.Stdin)
    for {
        displayMenu()
        choice, err := readUserChoice(reader)
        if err != nil {
            fmt.Println("Error reading choice:", err)
            continue
        }

        if choice == "0" {
            fmt.Println("\n[Exiting DirWizard CLI]")
            break
        }

        executeChoice(choice, reader)
    }
}

func readUserChoice(reader *bufio.Reader) (string, error) {
    fmt.Print("Enter your choice (or 0 to exit): ")
    choice, err := reader.ReadString('\n')
    if err != nil {
        return "", err
    }
    return strings.TrimSpace(choice), nil
}

func executeChoice(choice string, reader *bufio.Reader) {
    if choice == "7" {
        workingDir = getWorkingDirectory()
        clearScreen()
        displayHeader()
        return
    }
    if !runScript(choice) {
        fmt.Println("\nInvalid choice. Please try again.")
    }

    waitForContinue(reader)
}

func getWorkingDirectory() string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter the directory path to work on: ")
    dirPath, _ := reader.ReadString('\n')
    dirPath = strings.TrimSpace(dirPath)

    // Convert the provided directory path to an absolute path
    absDirPath, err := filepath.Abs(dirPath)
    if err != nil {
        fmt.Println("Error converting to absolute path:", err)
        return dirPath // fallback to the original input
    }
    return absDirPath
}

func runScript(choice string) bool {
    scriptName := getScriptName(choice)
    if scriptName == "" {
        return false
    }

    executeCommand("bash", scriptName, workingDir)
    return true
}

func getScriptName(choice string) string {
    switch choice {
    case "1":
        return "./pkg/scripts/rename_dir.sh"
    case "2":
        return "./pkg/scripts/check_dir.sh"
    case "3":
        return "./pkg/scripts/find_duplicate_dir.sh"
    case "4":
        return "./pkg/scripts/find_data_dir.sh"
    case "5":
        return "./pkg/scripts/mk_mock_dir.sh"
    case "6":
        return "./pkg/scripts/rm_mock_dir.sh"
    default:
        return ""
    }
}


func waitForContinue(reader *bufio.Reader) {
    fmt.Println("\nPress Enter to continue...")
    _, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input:", err)
    }
    clearScreen()
    displayHeader()
}
