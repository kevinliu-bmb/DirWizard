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

func main() {
	dirWizard := NewDirWizard()
	dirWizard.Run()
}

type DirWizard struct {
	workingDir string
	reader     *bufio.Reader
}

func NewDirWizard() *DirWizard {
	return &DirWizard{
		reader: bufio.NewReader(os.Stdin),
	}
}

func (dw *DirWizard) Run() {
	dw.clearScreen()
	var err error
	dw.workingDir, err = filepath.Abs(".")
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	dw.displayHeader()
	dw.processUserInput()
}

func (dw *DirWizard) clearScreen() {
	if runtime.GOOS == "windows" {
		dw.executeCommand("cmd", "/c", "cls")
	} else {
		dw.executeCommand("clear")
	}
}

func (dw *DirWizard) executeCommand(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command %s: %v\n", name, err)
	}
}

func (dw *DirWizard) displayHeader() {
	header := fmt.Sprintf(`=================================================
|         DirWizard - Directory Management      |
|         Version 0.1.0                         |
|         © 2024 Kevin Liu. All Rights Reserved |
=================================================
Working Directory: %s

`, filepath.Clean(dw.workingDir))
	fmt.Print(header)
}

func (dw *DirWizard) displayMenu() {
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

func (dw *DirWizard) processUserInput() {
	for {
		dw.displayMenu()
		choice, err := dw.readUserChoice()
		if err != nil {
			fmt.Println("Error reading choice:", err)
			continue
		}

		if choice == "0" {
			fmt.Println("\n[Exiting DirWizard CLI]")
			break
		}

		dw.executeChoice(choice)
	}
}

func (dw *DirWizard) readUserChoice() (string, error) {
	fmt.Print("Enter your choice (or 0 to exit): ")
	choice, err := dw.reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(choice), nil
}

func (dw *DirWizard) executeChoice(choice string) {
	if choice == "7" {
		dw.workingDir = dw.getWorkingDirectory()
		dw.clearScreen()
		dw.displayHeader()
		return
	}
	if !dw.runScript(choice) {
		fmt.Println("\nInvalid choice. Please try again.")
	}

	dw.waitForContinue()
}

func (dw *DirWizard) getWorkingDirectory() string {
	fmt.Print("Enter the directory path to work on: ")
	dirPath, _ := dw.reader.ReadString('\n')
	dirPath = strings.TrimSpace(dirPath)
	absDirPath, err := filepath.Abs(dirPath)
	if err != nil {
		fmt.Println("Error converting to absolute path:", err)
		return dirPath // fallback to the original input
	}
	return absDirPath
}

func (dw *DirWizard) runScript(choice string) bool {
	scriptName := dw.getScriptName(choice)
	if scriptName == "" {
		return false
	}
	dw.executeCommand("bash", scriptName, dw.workingDir)
	return true
}

func (dw *DirWizard) getScriptName(choice string) string {
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

func (dw *DirWizard) waitForContinue() {
	fmt.Println("\nPress Enter to continue...")
	if _, err := dw.reader.ReadString('\n'); err != nil {
		fmt.Println("Error reading input:", err)
	}
	dw.clearScreen()
	dw.displayHeader()
}
