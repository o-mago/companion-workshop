#!/bin/bash

# This script sets up the workshop environment:
#   1. Saves your Google Cloud Project ID to ~/project_id.txt
#   2. Saves your AI Studio API Key to ~/api_key.txt
#   3. Exports all required Google Cloud environment variables
#
# IMPORTANT: This script must be SOURCED so the exported variables
# are available in your current shell session.
# Usage: source ./setup.sh

# --- Configuration ---
PROJECT_FILE="$HOME/project_id.txt"
API_KEY_FILE="$HOME/api_key.txt"
GOOGLE_CLOUD_LOCATION="us-central1"
# ---------------------

# Using 'return' instead of 'exit' because this script must be sourced.
handle_error() {
  echo "Error: $1"
  return 1
}


# --- Part 1: Google Cloud Project ID ---
echo "--- Setting Google Cloud Project ID ---"

read -p "Please enter your Google Cloud project ID: " user_project_id

if [[ -z "$user_project_id" ]]; then
  handle_error "No project ID was entered."
  return 1
fi

echo "$user_project_id" > "$PROJECT_FILE" || { handle_error "Failed saving project ID."; return 1; }
echo "Successfully saved project ID."


# --- Part 2: AI Studio API Key ---
echo ""
echo "--- Setting AI Studio API Key ---"

read -p "Please enter your AI Studio API Key: " user_api_key

if [[ -z "$user_api_key" ]]; then
  handle_error "No API Key was entered."
  return 1
fi

echo "$user_api_key" > "$API_KEY_FILE" || { handle_error "Failed saving API Key."; return 1; }
echo "Successfully saved API Key."

export GOOGLE_API_KEY="$user_api_key"
echo "Exported GOOGLE_API_KEY."


# --- Part 3: Google Cloud Environment Variables ---
echo ""
echo "--- Setting Google Cloud Environment Variables ---"

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

export PROJECT_ID=$(gcloud config get project)
echo "Exported PROJECT_ID=$PROJECT_ID"

export PROJECT_NUMBER=$(gcloud projects describe "$PROJECT_ID" --format="value(projectNumber)")
echo "Exported PROJECT_NUMBER=$PROJECT_NUMBER"

export GOOGLE_CLOUD_PROJECT="$PROJECT_ID"
echo "Exported GOOGLE_CLOUD_PROJECT=$GOOGLE_CLOUD_PROJECT"

export GOOGLE_GENAI_USE_VERTEXAI="TRUE"
echo "Exported GOOGLE_GENAI_USE_VERTEXAI=$GOOGLE_GENAI_USE_VERTEXAI"

export GOOGLE_CLOUD_LOCATION="$GOOGLE_CLOUD_LOCATION"
echo "Exported GOOGLE_CLOUD_LOCATION=$GOOGLE_CLOUD_LOCATION"

export REGION="$GOOGLE_CLOUD_LOCATION"
echo "Exported REGION=$REGION"

source ~/.bashrc

echo ""
echo "--- Setup complete ---"
