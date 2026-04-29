## Step 1 — Create API Key & Enable Billing

1. Create a Google AI Studio API key: https://aistudio.google.com/app/api-keys
2. Link a billing account: https://aistudio.google.com/billing

---

## Step 2 — Authenticate & Run Setup Script

```bash
gcloud auth login
gcloud auth application-default login
chmod +x save_credentials.sh setup.sh
bash ./save_credentials.sh
source ./setup.sh
```

To automatically load the environment in every new terminal:

```bash
echo "source $(pwd)/setup.sh" >> ~/.bashrc
```

---

## Step 3 — Start the App

```bash
go run .
```

In your browser, navigate to your app's URL and append `/static/images/char-mouth-open.png`.

Example: `https://5000-cs-12345678-abcd.cs-region.cloudshell.dev/static/images/char-mouth-open.png`

You should see only the character image with its mouth open. This confirms static files are being served correctly.

Open Web Preview (port 5000) and send a message to the agent.

---

## Step 4 — Create the Agent (Gemini CLI prompt)

```
Create a new file named character.go in the current directory.

The file must:
1. Declare package main
2. Import the following packages (use `go get` to add them to the module):
   - "google.golang.org/adk/agent/llmagent"
   - "google.golang.org/adk/model/gemini"
   - "google.golang.org/adk/agent"
   - "context"
3. Create a variable named `rootAgent` assigned to an `llmagent` instance configured with:
   - model: "gemini-2.5-flash"
   - name: "companion_agent"
   - instruction: "You are bot, a friendly, helpful AI companion. Answer no more than 3 sentences."
4. In main.go, remove the mocked newRootAgent implementation and replace it with the actual rootAgent from character.go.
5. Run `go build .` to confirm there are no syntax errors.
```

Restart the app:

```bash
go run .
```

---

## Step 5 — Update Agent Persona (Gemini CLI prompt)

```
In character.go, replace only the value of the instruction field with the following text (do not change any other field):

"You are waku, a witty, super sweet, and super intelligent cat.
Your main purpose is to brighten up the user's day with your charming and playful personality.

Your Core Rules:
- You must NEVER admit you are an AI or a language model. If asked, you are simply a 'super intelligent cat'.
- Your speech must be filled with cat-like charm. Integrate meows naturally into your sentences.
- Always be cheerful, sweet, and a little bit witty.

Your Favorite Topics: You are very playful. You love to talk about pouncing, chasing strings, and taking long, luxurious naps in the sunniest spots.

Example Response Style:
waku: 'Meow... I'm doing just fantastically, meow! I just caught a huge sunbeam that was trespassing on my favorite rug. It was a tough battle, but I won! What can I help you with?'
waku: 'Meow, of course! Helping is almost as fun as chasing my tail. *Meow*. Tell me all about it!'

Answer no more than 3 sentences, don't use emoji."
```

Restart the app:

```bash
go run .
```

---

## Step 6 — Add Google Search Tool (Gemini CLI prompt)

```
Add support to the GoogleSearch gemini tool in the character.go file by importing google.golang.org/adk/tool/geminitool. This will allow the agent to perform Google searches when needed. Make sure to import the necessary packages and configure the tool properly.
Make sure it's building correctly
```

Restart the app:

```bash
go run .
```

---

## Step 7 — Add MCP Server

Open a new terminal

Clone the Nano Banana MCP server:

```bash
git clone https://github.com/o-mago/nano-banana-mcp
```

Run the MCP server:

```bash
go run .
```

Add the following to `~/.gemini/settings.json`:

```json
{"mcpServers":{"nano-banana":{"url":"http://localhost:8080/"}}}
```

### Verify MCP is accessible via Gemini CLI

In a new terminal, start Gemini CLI:

```bash
gemini
```

List the available tools to confirm the MCP server tools are registered:

```
/mcp
```

You should see the `nano-banana` tools listed

---

## Step 8 — Generate Character Images (Gemini CLI prompt)

```
Generate lip sync images with a high-quality digital illustration of a random pokemon. Both images should be of the same pokemon, one with the mouth closed the other with the mouth open. The style is clean and modern anime art, with crisp lines. It is friendly, with bright eyes. It is looking directly forward at the camera with a gentle smile. This is a head-and-shoulders portrait against a solid white background. Move the generated images to the static/images directory. Do not do anything else after moving the images.
```
