#!/bin/bash

# Prompts for Google Cloud Project ID and AI Studio API Key,
# then saves them to ~/project_id.txt and ~/api_key.txt.
# Usage: bash ./save_credentials.sh

PROJECT_FILE="$HOME/project_id.txt"
API_KEY_FILE="$HOME/api_key.txt"

echo "--- Setting Google Cloud Project ID ---"

read -p "Please enter your Google Cloud project ID: " user_project_id

if [[ -z "$user_project_id" ]]; then
  echo "Error: No project ID was entered."
  exit 1
fi

echo "$user_project_id" > "$PROJECT_FILE" || { echo "Error: Failed saving project ID."; exit 1; }
echo "Successfully saved project ID."


echo ""
echo "--- Setting AI Studio API Key ---"

read -p "Please enter your AI Studio API Key: " user_api_key

if [[ -z "$user_api_key" ]]; then
  echo "Error: No API Key was entered."
  exit 1
fi

echo "$user_api_key" > "$API_KEY_FILE" || { echo "Error: Failed saving API Key."; exit 1; }
echo "Successfully saved API Key."

echo ""
echo "--- Credentials saved. Run: source ./setup.sh ---"
