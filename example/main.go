package main

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	cloudtrace "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
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

func initTracer() (func(), error) {
	projectID := os.Getenv("PROJECT_ID")

	exporter, err := cloudtrace.New(cloudtrace.WithProjectID(projectID))
	if err != nil {
		return nil, err
	}

	res, err := resource.New(context.Background(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(appName),
		),
	)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(tp)

	return func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("failed to shutdown tracer provider: %v", err)
		}
	}, nil
}

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

	shutdown, err := initTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer shutdown()

	rootAgent, err := newRootAgent(ctx)
	if err != nil {
		log.Fatal("failed to create agent:", err)
	}

	if rootAgent != nil {
		sessionSvc := session.InMemoryService()

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
