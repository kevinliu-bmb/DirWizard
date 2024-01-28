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
LOG_FILE="${LOG_DIR}/data_directories_log_${TIMESTAMP}.txt"

# Create log directory if it doesn't exist
mkdir -p "$LOG_DIR"

# Define the file extensions that you consider as 'raw data'.
declare -a DATA_EXTENSIONS=("csv" "xlsx" "txt" "dat" "fasta" "fastq" "vcf" "bam" "sam" "gff" "gff3" "bed" "ab1" "mzXML" "mzML" "raw" "dta" "mgf" "cel" "gpr" "fid" "nmr" "pdb" "cif" "dicom" "tiff" "tif" "nii" "svs" "abf" "edf" "pdbqt" "mol2" "sdf" "json" "xml" "hl7" "fhir" "gz")

# Function to search for data files and log directories
function find_data_files() {
    for ext in "${DATA_EXTENSIONS[@]}"; do
        # Find directories containing files with the specified extension, count them, and append to the log file
        find "$BASE_DIR" -type f -name "*.$ext" | awk -F/ '{OFS=FS; NF--; print}' | sort | uniq -c | awk -v ext="$ext" '{print ext " (" $1 " files): " $2}' >> "$LOG_FILE"
    done
}

# Execute the function
find_data_files

# Remove duplicate entries in the log file
sort -u "$LOG_FILE" -o "$LOG_FILE"

echo
echo "[Data directory search completed]"
