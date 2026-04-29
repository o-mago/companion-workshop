#!/bin/bash

# Reads saved credentials and exports all required environment variables
# into the current shell session.
#
# IMPORTANT: Must be SOURCED so exports persist in your shell.
# Usage: source ./setup.sh

PROJECT_FILE="$HOME/project_id.txt"
API_KEY_FILE="$HOME/api_key.txt"

handle_error() {
  echo "Error: $1"
  return 1
}

if [[ ! -f "$PROJECT_FILE" ]]; then
  handle_error "Project ID file not found. Run: bash ./save_credentials.sh"
  return 1
fi

if [[ ! -f "$API_KEY_FILE" ]]; then
  handle_error "API Key file not found. Run: bash ./save_credentials.sh"
  return 1
fi

user_project_id=$(cat "$PROJECT_FILE")
user_api_key=$(cat "$API_KEY_FILE")

echo "--- Setting Environment Variables ---"

echo "Checking gcloud authentication status..."
if gcloud auth print-access-token > /dev/null 2>&1; then
  echo "gcloud is authenticated."
else
  echo "Error: gcloud is not authenticated."
  echo "Please log in by running: gcloud auth login"
  return 1
fi

echo "Setting gcloud config project to: $user_project_id"
gcloud config set project "$user_project_id" --quiet

export GOOGLE_API_KEY="$user_api_key"
echo "Exported GOOGLE_API_KEY."

export PROJECT_ID=$(gcloud config get project)
echo "Exported PROJECT_ID=$PROJECT_ID"

echo ""
echo "--- Setup complete ---"
