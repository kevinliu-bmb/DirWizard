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
LOG_DIR="${BASE_DIR}/dirwizard_log"
TIMESTAMP=$(date +"%Y-%m-%d_%H-%M-%S")
LOG_FILE="${LOG_DIR}/rename_log_${TIMESTAMP}.txt"

# Create log directory if it doesn't exist
mkdir -p "$LOG_DIR"

function rename_dirs() {
    find "$1" -depth -type d ! -name .git ! -name __pycache__ | while IFS= read -r dir; do
        local base_dir=$(dirname "$dir")
        local dir_name=$(basename "$dir")

        # Transform the directory name
        local new_dir_name=$(echo "$dir_name" | tr '[:upper:]' '[:lower:]' | tr -s ' -' '_' | sed 's/^_\+//; s/_\+$//')

        # If the transformed name is different, ask the user if directory should be renamed
        if [[ "$dir_name" != "$new_dir_name" ]]; then
            while true; do
                echo -n "Rename directory $dir_name to $new_dir_name? (y/n): "
                read -r answer < /dev/tty
                case $answer in
                    [Yy]* )
                        echo "Renaming directory $dir_name to $new_dir_name"
                        local new_dir_path="${base_dir}/${new_dir_name}"
                        mv "$dir" "$new_dir_path" 2>>"$LOG_FILE" && echo "Renamed: $dir -> $new_dir_path" >> "$LOG_FILE"
                        break
                        ;;
                    [Nn]* )
                        echo "Skipping $dir_name"
                        break
                        ;;
                    * )
                        echo "Please answer yes (y) or no (n)."
                        ;;
                esac
            done
        fi
    done
}

rename_dirs "$BASE_DIR"
echo
echo "[Rename directory processed successfully]"
