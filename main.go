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
	"runtime"
	"strings"
)

func main() {
    clearScreen()
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
    header := `=================================================
|         DirWizard - Directory Management      |
|         Version 0.1.0                         |
|         © 2024 Kevin Liu. All Rights Reserved |
=================================================
Welcome to DirWizard, your go-to tool for efficient
directory management and data organization.

`
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
    if !runScript(choice) {
        fmt.Println("\nInvalid choice. Please try again.")
    }

    waitForContinue(reader)
}

func runScript(choice string) bool {
    scriptName := getScriptName(choice)
    if scriptName == "" {
        return false
    }

    executeCommand("bash", scriptName)
    return true
}

func getScriptName(choice string) string {
    switch choice {
    case "1":
        return "./src/rename_directories.sh"
    case "2":
        return "./src/check_compliance.sh"
    case "3":
        return "./src/find_duplicates.sh"
    case "4":
        return "./src/search_data_paths.sh"
    case "5":
        return "./src/generate_mock_structure.sh"
    case "6":
        return "./src/clear_mock_structure.sh"
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
