from flask import Flask, jsonify
import datetime

app = Flask(__name__)

@app.route('/health', methods=['GET'])
def health_check():
    return jsonify({
        "status": "UP",
        "timestamp": datetime.datetime.now().isoformat(),
        "location": "Kaduna-Axis-DC",
        "system": "Kubernetes-WSL"
    }), 200

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
