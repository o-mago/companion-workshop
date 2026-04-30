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
   - instruction: "You are Gophi, a friendly Go gopher. Answer no more than 3 sentences."
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
{"mcpServers":{"nano-banana":{"url":"http://localhost:8090/"}}}
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
Generate lip sync images of Gophi, the Go gopher mascot. Both images should be of the same character, one with the mouth closed and one with the mouth open. The style is a high-quality digital illustration: clean, friendly, slightly chubby blue gopher with big bright eyes, wearing a tiny Botafogo football team jersey, looking directly forward at the camera. Head-and-shoulders portrait against a solid white background. Move the generated images to the static/images directory. Do not do anything else after moving the images.
```

Run again the app:

```bash
go run .
```

---

## Step 9 — Send Agent Traces to GCP Cloud Trace (Gemini CLI prompt)

```
Add OpenTelemetry tracing to the application so that agent traces are exported to GCP Cloud Trace.

Follow these steps exactly:

1. Add the Cloud Trace exporter package:
   go get github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace

2. In main.go, add a new function called initTracer that:
   - Imports:
       cloudtrace "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
       "go.opentelemetry.io/otel"
       "go.opentelemetry.io/otel/sdk/resource"
       sdktrace "go.opentelemetry.io/otel/sdk/trace"
       semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
   - Creates a Cloud Trace exporter using cloudtrace.New() with no options (it will use ADC automatically)
   - Creates a TracerProvider with:
       - The Cloud Trace exporter as a BatchSpanProcessor
       - A Resource with service.name set to the appName constant
   - Registers the provider globally with otel.SetTracerProvider
   - Returns a shutdown function (func()) and an error

3. In the main() function, call initTracer() right after creating the context, before anything else. If it returns an error, log.Fatal it. Defer the shutdown function.

4. Run go build . to confirm there are no syntax errors.
```

Restart the app:

```bash
go run .
```

Send a few messages in the chat, then open the GCP Cloud Trace Explorer to see the traces:

```
https://console.cloud.google.com/traces/list?project=YOUR_PROJECT_ID
```

> **Note:** Traces may take up to 30 seconds to appear in the console. Make sure `PROJECT_ID` is set in your environment (`echo $PROJECT_ID`).
