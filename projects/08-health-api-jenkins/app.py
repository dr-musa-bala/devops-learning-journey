from flask import Flask
from datetime import datetime
import os

app = Flask(__name__)

@app.route('/health')
def health_check():
    # Get environment variable (default if not set)
    location = os.getenv("APP_LOCATION", "Unknown-Location")
    
    return {
        "status": "UP",
        "location": location,
        "system": "Kubernetes-WSL",
        "timestamp": datetime.now().isoformat()
    }

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=80)
