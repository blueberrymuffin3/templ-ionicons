#!/usr/bin/env bash

set -euo pipefail

# Set the source and destination directories
source_dir="ionicons/src/svg/"
destination_dir="data/"

# Create the destination directory if it doesn't exist
mkdir -p "$destination_dir"

# Iterate through all files in the source directory
for file in "$source_dir"*.svg; do
    # Get the filename without the path
    filename=$(basename "$file")

    # Remove the specified substrings from the SVG file and save it to the destination directory
    sed -e 's|<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">||' -e 's|</svg>||' "$file" > "$destination_dir$filename"
done
