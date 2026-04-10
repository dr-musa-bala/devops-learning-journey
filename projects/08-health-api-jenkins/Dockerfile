FROM python:3.9-slim

WORKDIR /app

# --- This is the part that handles the new library ---
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
# -----------------------------------------------------

COPY . .

# Your existing lines:
EXPOSE 80
CMD ["python", "app.py"]
