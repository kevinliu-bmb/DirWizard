#!/bin/bash

# This file is part of DirWizard.
# Copyright (C) 2024 Kevin Liu

# DirWizard is free software; you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation; either version 2 of the License, or
# (at your option) any later version.

# DirWizard is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.

# You should have received a copy of the GNU General Public License
# along with DirWizard.  If not, see <http://www.gnu.org/licenses/>.

# Use the first argument as BASE_DIR, default to current directory if not provided
BASE_DIR="${1:-.}"
MOCK_DIR="${BASE_DIR}/mock_dir"
LOG_DIR="${BASE_DIR}/dirwizard_log"
TIMESTAMP=$(date +"%Y-%m-%d_%H-%M-%S")
LOG_FILE="${LOG_DIR}/mock_generation_log_${TIMESTAMP}.txt"

# Create log and mock directories if they don't exist
mkdir -p "$LOG_DIR"
mkdir -p "$MOCK_DIR"

# Function to create mock directories and files
function create_mock_structure() {
    # (Include the directory and file creation commands here as previously provided)
    echo "Mock file structure created at $MOCK_DIR" >> "$LOG_FILE"

    mkdir -p "$MOCK_DIR"

    # Create a variety of directories and subdirectories
    mkdir -p "$MOCK_DIR/Directory One"
    mkdir -p "$MOCK_DIR/directory_two"
    mkdir -p "$MOCK_DIR/Directory-Three"
    mkdir -p "$MOCK_DIR/Directory_Four"
    mkdir -p "$MOCK_DIR/directoryFive"
    mkdir -p "$MOCK_DIR/directory_Six-Seven"
    mkdir -p "$MOCK_DIR/Directory Eight/Nested One"
    mkdir -p "$MOCK_DIR/Directory Eight/Nested_Two"
    mkdir -p "$MOCK_DIR/__pycache__"
    mkdir -p "$MOCK_DIR/.git"

    # Create duplicate directories
    mkdir -p "$MOCK_DIR/DuplicateDir/Child1"
    mkdir -p "$MOCK_DIR/DuplicateDir/Child2"
    mkdir -p "$MOCK_DIR/DuplicateDirCopy/Child1"
    mkdir -p "$MOCK_DIR/DuplicateDirCopy/Child2"

    # Create some files in directories
    touch "$MOCK_DIR/Directory One/file1.txt"
    touch "$MOCK_DIR/directory_two/file2.txt"
    touch "$MOCK_DIR/Directory-Three/file3.txt"
    touch "$MOCK_DIR/Directory_Four/file4.txt"
    touch "$MOCK_DIR/directoryFive/file5.txt"
    touch "$MOCK_DIR/directory_Six-Seven/file6.txt"
    touch "$MOCK_DIR/Directory Eight/Nested One/file7.txt"
    touch "$MOCK_DIR/Directory Eight/Nested_Two/file8.txt"
    touch "$MOCK_DIR/__pycache__/file9.txt"
    touch "$MOCK_DIR/.git/file10.txt"
    touch "$MOCK_DIR/DuplicateDir/Child1/file11.txt"
    touch "$MOCK_DIR/DuplicateDir/Child2/file12.txt"
    touch "$MOCK_DIR/DuplicateDirCopy/Child1/file11.txt"
    touch "$MOCK_DIR/DuplicateDirCopy/Child2/file12.txt"

    echo
    echo "Mock file structure created at $MOCK_DIR"
}

create_mock_structure
