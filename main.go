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
	"strings"
)

func main() {
    clearScreen()
    displayHeader()
    reader := bufio.NewReader(os.Stdin)

    for {
        displayMenu()
        fmt.Print("Enter your choice (or 0 to exit): ")
        choices, _ := reader.ReadString('\n')
        choices = strings.TrimSpace(choices)

        if choices == "0" {
            fmt.Println("[Exiting DirWizard CLI]")
            break
        }

        executeChoice(choices)
    }
}

func clearScreen() {
    cmd := exec.Command("clear") // For Windows, use exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func displayHeader() {
    fmt.Println("================================")
    fmt.Println("|     DirWizard CLI            |")
    fmt.Println("|     Developed by Kevin Liu   |")
    fmt.Println("================================")
}

func displayMenu() {
    fmt.Println()
    fmt.Println("Menu:")
    fmt.Println("1 - Rename Directories")
    fmt.Println("2 - Check Compliance")
    fmt.Println("3 - Find Duplicates")
    fmt.Println()
    fmt.Println("4 - [dev test] Generate Mock File Structure")
    fmt.Println("5 - [dev test] Clear Mock File Structure")
    fmt.Println()
    fmt.Println("0 - Exit")
    fmt.Println()
}

func executeChoice(choice string) {
    switch choice {
    case "1":
        runScript("./src/rename_directories.sh")
    case "2":
        runScript("./src/check_compliance.sh")
    case "3":
        runScript("./src/find_duplicates.sh")
    case "4":
        runScript("./src/generate_mock_structure.sh")
    case "5":
        runScript("./src/clear_mock_structure.sh")
    default:
        fmt.Println("Invalid choice. Please try again.")
    }
}

func runScript(scriptName string) {
    cmd := exec.Command("bash", scriptName)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        fmt.Println("Error executing script:", scriptName, "-", err)
    }
}
