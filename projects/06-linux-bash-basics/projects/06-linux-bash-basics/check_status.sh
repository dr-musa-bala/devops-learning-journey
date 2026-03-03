#!/bin/bash

# Configuration
URL=${API_URL:-"http://localhost:8080/health"}
MAX_ATTEMPTS=5
SLEEP_TIME=2

echo "---------------------------------------"
echo "🔍 DEVOPS SYSTEM CHECK IN PROGRESS..."
echo "Target: $URL"
echo "---------------------------------------"

for ((i=1; i<=MAX_ATTEMPTS; i++)); do
    # -s: silent, -o: ignore body, -w: return only status code
    STATUS=$(curl -s -o /dev/null -w "%{http_code}" "$URL")

    if [ "$STATUS" == "200" ]; then
        echo "✅ SUCCESS: Your Go API is UP (Attempt $i)."
        echo "---------------------------------------"
        exit 0
    fi

    echo "⏳ Attempt $i/$MAX_ATTEMPTS: API not ready (Status: $STATUS). Retrying in ${SLEEP_TIME}s..."
    sleep $SLEEP_TIME
done

echo "❌ ALERT: The Go API is DOWN after $MAX_ATTEMPTS attempts."
echo "---------------------------------------"
exit 1
