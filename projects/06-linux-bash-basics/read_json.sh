#!/bin/bash
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
