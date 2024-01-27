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
    reader := bufio.NewReader(os.Stdin)

    for {
        displayMenu()
        fmt.Print("Enter your choice (or 0 to exit): ")
        choice, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading choice:", err)
            continue
        }
        choice = strings.TrimSpace(choice)

        if choice == "0" {
            fmt.Println("\n[Exiting DirWizard CLI]")
            break
        }

        executeChoice(choice, reader)
    }
}

func clearScreen() {
    var cmd *exec.Cmd
    if runtime.GOOS == "windows" {
        cmd = exec.Command("cmd", "/c", "cls")
    } else {
        cmd = exec.Command("clear")
    }
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func displayHeader() {
    fmt.Println("=================================================")
    fmt.Println("|         DirWizard - Directory Management      |")
    fmt.Println("|         Version 0.1.0                         |")
    fmt.Println("|         © 2024 Kevin Liu. All Rights Reserved |")
    fmt.Println("=================================================")
    fmt.Println("Welcome to DirWizard, your go-to tool for efficient")
    fmt.Println("directory management and data organization.")
    fmt.Println()
}

func displayMenu() {
    fmt.Println("Menu:")
    fmt.Println("1 - Rename Directories")
    fmt.Println("2 - Check Compliance")
    fmt.Println("3 - Find Duplicates")
    fmt.Println("4 - Search Data Paths")
    fmt.Println()
    fmt.Println("5 - [dev test] Generate Mock File Structure")
    fmt.Println("6 - [dev test] Clear Mock File Structure")
    fmt.Println()
    fmt.Println("0 - Exit")
    fmt.Println()
}

func executeChoice(choice string, reader *bufio.Reader) {
    if !runScript(choice) {
        fmt.Println("\nInvalid choice. Please try again.")
    }

    fmt.Println("\nPress Enter to continue...")
    _, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input:", err)
    }
    clearScreen()
    displayHeader()
}

func runScript(choice string) bool {
    var scriptName string
    switch choice {
    case "1":
        scriptName = "./src/rename_directories.sh"
    case "2":
        scriptName = "./src/check_compliance.sh"
    case "3":
        scriptName = "./src/find_duplicates.sh"
    case "4":
        scriptName = "./src/search_data_paths.sh"
    case "5":
        scriptName = "./src/generate_mock_structure.sh"
    case "6":
        scriptName = "./src/clear_mock_structure.sh"
    default:
        return false
    }

    cmd := exec.Command("bash", scriptName)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        fmt.Println("Error executing script:", scriptName, "-", err)
        return false
    }
    return true
}
