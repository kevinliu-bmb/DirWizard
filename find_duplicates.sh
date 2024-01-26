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

BASE_DIR=".."
LOG_DIR="./log"
TIMESTAMP=$(date +"%Y-%m-%d_%H-%M-%S")
LOG_FILE="${LOG_DIR}/duplicates_log_${TIMESTAMP}.txt"
TEMP_FILE="${LOG_DIR}/temp_signatures_${TIMESTAMP}.txt"

# Create log and temporary directories if they don't exist
mkdir -p "$LOG_DIR"

function create_signature() {
    local dir=$1
    find "$dir" -type f -exec basename {} \; | sort | shasum | awk '{print $1}'
}

# Clear or create the temporary file for signatures
> "$TEMP_FILE"

function find_duplicates() {
    while IFS= read -r -d '' dir; do
        if [[ "$dir" != "$BASE_DIR" ]]; then
            local sig=$(create_signature "$dir")
            if grep -q "$sig" "$TEMP_FILE"; then
                local original_dir=$(grep "$sig" "$TEMP_FILE" | cut -d':' -f1)
                echo "Duplicate: $dir AND $original_dir" >> "$LOG_FILE"
            else
                echo "$dir:$sig" >> "$TEMP_FILE"
            fi
        fi
    done < <(find "$BASE_DIR" -type d -print0)
}

find_duplicates

# Clean up
rm "$TEMP_FILE"
