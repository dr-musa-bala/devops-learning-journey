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

WINDOWS_HOST_IP=$(ip route show | grep default | awk '{print $3}')
API_URL="http://$WINDOWS_HOST_IP:8080/health"

echo "📡 Fetching JSON Data from Windows Host ($WINDOWS_HOST_IP)..."

# 1. Capture the full JSON response into a variable
RESPONSE=$(curl -s "$API_URL")

# 2. Use 'jq' to extract specific fields
# The '.' means the root of the JSON object
STATUS=$(echo "$RESPONSE" | jq -r '.status')
MESSAGE=$(echo "$RESPONSE" | jq -r '.message')

echo "---------------------------------------"
echo "API STATUS  : $STATUS"
echo "API MESSAGE : $MESSAGE"
echo "---------------------------------------"

if [ "$STATUS" == "Healthy" ]; then
    echo "✅ System is verified by JSON data."
else
    echo "⚠️ JSON Status unexpected: $STATUS"
fi

