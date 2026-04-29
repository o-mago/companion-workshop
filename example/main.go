package main

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/adk/agent"
	"google.golang.org/adk/runner"
	"google.golang.org/adk/session"
	"google.golang.org/genai"
)

var (
	agentRunner *runner.Runner
	tmpl        *template.Template
)

const (
	appName   = "companion"
	userID    = "inapp_user"
	sessionID = "default_session"
)

type chatRequest struct {
	Message string `json:"message"`
}

type chatResponse struct {
	Response string `json:"response"`
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleChat(w http.ResponseWriter, r *http.Request) {
	var req chatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if agentRunner == nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(chatResponse{Response: req.Message})

		return
	}

	msg := genai.NewContentFromText(req.Message, genai.RoleUser)

	var sb strings.Builder
	for event, err := range agentRunner.Run(r.Context(), userID, sessionID, msg, agent.RunConfig{}) {
		if err != nil {
			log.Printf("run error: %v", err)
			continue
		}
		if event.IsFinalResponse() && event.Content != nil {
			for _, p := range event.Content.Parts {
				sb.WriteString(p.Text)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chatResponse{Response: sb.String()})
}

func main() {
	ctx := context.Background()
	projectID := os.Getenv("PROJECT_ID")

	// 1. Create the exporter
	exporter, err := texporter.New(texporter.WithProjectID(projectID))
	if err != nil {
		log.Fatalf("failed to create exporter: %v", err)
	}

	// 2. Install the Trace Provider
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
	)
	defer tp.Shutdown(ctx)
	otel.SetTracerProvider(tp)

	rootAgent, err := NewRootAgent(context.Background())
	if err != nil {
		log.Printf("failed to create root agent: %v", err)
		log.Println("Server will run without AI agent functionality.")
	}

	if rootAgent != nil {
		sessionSvc := session.InMemoryService()

		var err error
		agentRunner, err = runner.New(runner.Config{
			AppName:           appName,
			Agent:             rootAgent,
			SessionService:    sessionSvc,
			AutoCreateSession: true,
		})
		if err != nil {
			log.Fatal("failed to create runner:", err)
		}
	}

	tmpl = template.Must(template.ParseFiles("templates/index.html"))

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handleIndex)
	mux.HandleFunc("POST /chat", handleChat)
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server listening on :5000")
	log.Fatal(http.ListenAndServe(":5000", mux))
}
