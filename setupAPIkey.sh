#!/bin/bash

# --- Function for error handling ---
handle_error() {
  echo "Error: $1"
  exit 1
}

# --- Part 1: Set Google Cloud Project ID ---
API_KEY_FILE="$HOME/api_key.txt"
echo "--- Setting AI Studio API Key File ---"

read -p "Please enter your AI Studio API Key: " user_api_key

if [[ -z "$user_api_key" ]]; then
  handle_error "No API KEY was entered."
fi

echo "You entered: $user_api_key"
echo "$user_api_key" > "$API_KEY_FILE"

if [[ $? -ne 0 ]]; then
  handle_error "Failed saving your API KEY: $user_api_key."
fi
echo "Successfully saved API KEY."

# --- Part 2: Set GOOGLE_API_KEY environment variable ---
echo "--- Setting GOOGLE_API_KEY environment variable ---"

SHELL_RC=""
if [[ "$SHELL" == *"zsh"* ]]; then
  SHELL_RC="$HOME/.zshrc"
elif [[ "$SHELL" == *"bash"* ]]; then
  SHELL_RC="$HOME/.bashrc"
else
  handle_error "Unsupported shell: $SHELL. Please set GOOGLE_API_KEY manually."
fi

if grep -q "export GOOGLE_API_KEY=" "$SHELL_RC" 2>/dev/null; then
  sed -i.bak "s|export GOOGLE_API_KEY=.*|export GOOGLE_API_KEY=\"$user_api_key\"|" "$SHELL_RC"
  echo "Updated GOOGLE_API_KEY in $SHELL_RC."
else
  echo "export GOOGLE_API_KEY=\"$user_api_key\"" >> "$SHELL_RC"
  echo "Added GOOGLE_API_KEY to $SHELL_RC."
fi

export GOOGLE_API_KEY="$user_api_key"
echo "GOOGLE_API_KEY is set for the current session."

echo "--- Setup complete ---"
exit 0