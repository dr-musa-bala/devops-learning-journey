#!/bin/bash

# Check if a filename was provided
if [ -z "$1" ]; then
    echo "Usage: $0 <json_file>"
    exit 1
fi

# Parse the 'status' field from the JSON file
# Note: We use quotes around "$1" to satisfy ShellCheck (SC2086)
STATUS=$(jq -r '.status' "$1")

echo "The system status is: $STATUS"
# end of script
