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
LOG_FILE="${LOG_DIR}/compliance_log_${TIMESTAMP}.txt"

# Create log directory if it doesn't exist
mkdir -p "$LOG_DIR"

function check_compliance() {
    find "$1" -type d ! -name .git ! -name __pycache__ | while read -r dir; do
        if [[ ! $(basename "$dir") =~ ^[a-z0-9_]+$ ]]; then
            echo "Non-compliant: $dir" >> "$LOG_FILE"
        fi
    done
}

check_compliance "$BASE_DIR"
