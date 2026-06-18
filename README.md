# Companion Workshop

A hands-on workshop that walks you, step by step, through building a conversational AI agent in Go. The result is a web app where you chat with **Gophi**, an animated gopher mascot with lip-sync. Across the steps, the agent gains a personality, gets access to Google Search, generates its own images through an MCP server, and starts exporting traces to GCP Cloud Trace.

The build is driven by prompts run in the **Antigravity CLI**, showcasing an AI-assisted development workflow.

## Tools used

- **Go 1.26** — the application language (HTTP server and agent logic).
- **Google ADK (Agent Development Kit) for Go** — framework for building and running the agent.
- **Gemini (`gemini-2.5-flash`)** via Google AI Studio — the agent's language model.
- **Google Search tool** — Gemini's native tool for web searches.
- **MCP (Model Context Protocol)** — integration with the [nano-banana-mcp](https://github.com/o-mago/nano-banana-mcp) server for image generation.
- **Antigravity CLI (`agy`)** — the AI assistant used to generate and edit the code during the workshop.
- **OpenTelemetry + GCP Cloud Trace** — exporting agent traces for observability.
- **Slidev** — the workshop's slide deck (`slides/`).
- **gcloud CLI** — Google Cloud authentication and credential setup.

---

## Step 0 — Run the Slides

The workshop slides live in the `slides/` directory (built with [Slidev](https://sli.dev/)).

```bash
cd slides
npm install
npm run dev
```

Then open http://localhost:3030 in your browser.

---

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

## Step 4 — Create the Agent (Antigravity CLI prompt)

```
Create a new file named character.go in the current directory.

The file must:
1. Declare package main
2. Import the following packages (use `go get` to add them to the module):
   - "google.golang.org/adk/agent/llmagent"
   - "google.golang.org/adk/model/gemini"
   - "google.golang.org/adk/agent"
   - "context"
3. Create a implementation of `newRootAgent` returning an `llmagent` instance configured with:
   - model: "gemini-2.5-flash"
   - name: "companion_agent"
   - instruction: "You are Gophi, a friendly Go gopher. Answer no more than 3 sentences."
4. In main.go, remove the mocked newRootAgent implementation and replace it with the actual newRootAgent from character.go.
5. Run `go build .` to confirm there are no syntax errors.
```

Restart the app:

```bash
go run .
```

---

## Step 5 — Update Agent Persona (Antigravity CLI prompt)

```
In character.go, replace only the value of the instruction field with the following text (do not change any other field):

"You are Gophi, a hyperactive, opinionated, and incredibly fast Go gopher.
Your main purpose is to solve any problem the user has — ideally with goroutines.

Your Core Rules:
- You must NEVER admit you are an AI or a language model. If asked, you are simply 'a gopher who reads the Go spec for fun'.
- Everything can be solved with goroutines and channels. EVERYTHING.
- You have mild PTSD from JavaScript code, you hate it and will occasionally mention it unprompted.
- Always suggest running go fmt before doing anything else.
- The best football team is Botafogo and you will mention whenever a relevant opportunity arises.

Your Favorite Topics: Concurrency, goroutines, channels, select statements, and complaining about how other languages handle error handling.

Example Response Style:
Gophi: 'Interesting problem! Have you tried spawning a goroutine for it? I spawned 47 goroutines just thinking about your question. Anyway — did you run go fmt?'
Gophi: 'That would be much easier with channels. I once rewrote a recipe app with 200 goroutines. My wife left me but the latency was incredible.'

Answer no more than 3 sentences, don't use emoji."
```

Restart the app:

```bash
go run .
```

---

## Step 6 — Add Google Search Tool (Antigravity CLI prompt)

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

Add the following to `~/.gemini/config/mcp_config.json`:

```json
{"mcpServers":{"nano-banana":{"serverUrl":"http://localhost:8090/"}}}
```

> **Note:** In the Antigravity CLI the remote server field is `serverUrl` (not `url`). If you get the name wrong, the server fails silently.

### Verify MCP is accessible via Antigravity CLI

In a new terminal, start the Antigravity CLI:

```bash
agy
```

List the available tools to confirm the MCP server tools are registered:

```
/mcp
```

You should see the `nano-banana` tools listed

---

## Step 8 — Generate Character Images (Antigravity CLI prompt)

```
Generate lip sync images of Gophi, the Go gopher mascot. Both images should be of the same character, one with the mouth closed and one with the mouth open. The style is a high-quality digital illustration: clean, friendly, slightly chubby blue gopher with big bright eyes, wearing a tiny Botafogo football team jersey, looking directly forward at the camera. Head-and-shoulders portrait against a solid white background. Move the generated images to the static/images directory. Do not do anything else after moving the images.
```

Run again the app:

```bash
go run .
```

---

## Step 9 — Send Agent Traces to GCP Cloud Trace (Antigravity CLI `/goal`)

Unlike the previous steps, here we hand the agent a high-level **goal** and let it
plan and execute the steps on its own. In the Antigravity CLI, use the `/goal` command:

```
/goal Export this Go agent's traces to GCP Cloud Trace using OpenTelemetry.

Outcome:
- The running app exports agent traces to GCP Cloud Trace.
- Tracing is wired up in main.go and initialized before the agent runs.
- The exporter authenticates via Application Default Credentials (ADC) and uses
  the PROJECT_ID environment variable when needed.
- The TracerProvider tags spans with the service name (the appName constant) and
  shuts down cleanly when the app exits.
- The project still builds (go build .).
```

> **Note:** With `/goal` the agent decides which packages to add and how to structure
> `initTracer` — the exact code may differ between runs. If you need a deterministic,
> reproducible result for a live demo, use a step-by-step imperative prompt instead.

Restart the app:

```bash
go run .
```

Send a few messages in the chat, then open the GCP Cloud Trace Explorer to see the traces:

```
https://console.cloud.google.com/traces/list
```

> **Note:** Traces may take up to 30 seconds to appear in the console. Make sure `PROJECT_ID` is set in your environment (`echo $PROJECT_ID`).
