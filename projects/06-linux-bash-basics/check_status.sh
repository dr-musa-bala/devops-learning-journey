#!/bin/bash

# Configuration: The address of your Go API
API_URL="http://localhost:8080/health"

echo "---------------------------------------"
echo "🔍 DEVOPS SYSTEM CHECK IN PROGRESS..."
echo "---------------------------------------"

# 1. 'curl' knocks on the door of your Go API
# 2. 'grep' looks for the "200" success code in the response
# 3. '> /dev/null' hides the messy technical details from the screen
if [ $status == "UP"
    echo "✅ SUCCESS: Your Go API is UP and running."
    echo "Report: System Healthy."
else
    echo "❌ ALERT: The Go API is DOWN."
    echo "Action Required: Start your Go server in your other terminal!"
fi

echo "---------------------------------------"
echo "Check complete at: $(date)"
