from flask import Flask, jsonify
import os
from datetime import datetime
from prometheus_flask_exporter import PrometheusMetrics  # New library!

app = Flask(__name__)
metrics = PrometheusMetrics(app)  # This automatically tracks request numbers and latency

@app.route('/health')
def health():
    return jsonify({
        "status": "UP",
        "location": os.getenv("APP_LOCATION", "Kaduna-Axis-DC"),
        "system": "Kubernetes-WSL",
        "timestamp": datetime.now().isoformat()
    })

if __name__ == '__main__':
    # Running on Port 80 as per our Phase 10 standards
    app.run(host='0.0.0.0', port=80)
