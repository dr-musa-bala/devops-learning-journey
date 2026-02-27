#!/bin/bash

# Use the API_URL variable if provided, otherwise default to localhost
# This makes the script flexible for both local and Docker use!
TARGET_URL=${API_URL:-"http://localhost:8080/health"}

echo "---------------------------------------"
echo "🔍 DEVOPS SYSTEM CHECK IN PROGRESS..."
echo "Target: $TARGET_URL"
echo "---------------------------------------"

# Use a standard GET and check for the word "Healthy"
if curl -s "$TARGET_URL" | grep -q "Healthy"; then
    echo "✅ SUCCESS: Your Go API is UP."
else
    echo "❌ ALERT: The Go API is DOWN."
fi

echo "---------------------------------------"
echo "Check complete at: $(date)"
