//go:build ignore

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/genai"
)

const model = "gemini-2.5-flash-image"

var cfg = &genai.GenerateContentConfig{
	ResponseModalities: []string{"IMAGE", "TEXT"},
}

const basePrompt = `A high-quality digital illustration of Gophi, a Go gopher mascot.
Clean, friendly, slightly chubby blue gopher with big bright eyes, wearing a tiny Go t-shirt (white tee with the Go gopher logo on it).
Looking directly forward at the camera with a cheerful expression.
Head-and-shoulders portrait against a solid white background.
Style: clean vector-like digital art, crisp outlines, vibrant colors.`

func extractImage(resp *genai.GenerateContentResponse) ([]byte, error) {
	if len(resp.Candidates) == 0 {
		return nil, fmt.Errorf("no candidates")
	}
	for _, part := range resp.Candidates[0].Content.Parts {
		if part.InlineData != nil {
			return part.InlineData.Data, nil
		}
	}
	return nil, fmt.Errorf("no image data in response")
}

func main() {
	ctx := context.Background()

	if os.Getenv("GOOGLE_API_KEY") == "" {
		log.Fatal("GOOGLE_API_KEY not set")
	}

	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	outDir := "static/images"

	// Generate mouth-open image
	fmt.Println("Generating mouth-open gopher...")
	respOpen, err := client.Models.GenerateContent(ctx, model,
		genai.Text(basePrompt+"\nThe gopher's mouth is wide open, mid-sentence, as if enthusiastically explaining goroutines."),
		cfg,
	)
	if err != nil {
		log.Fatal("mouth-open:", err)
	}
	openData, err := extractImage(respOpen)
	if err != nil {
		log.Fatal("mouth-open extract:", err)
	}
	if err := os.WriteFile(outDir+"/char-mouth-open.png", openData, 0644); err != nil {
		log.Fatal(err)
	}
	fmt.Println("✓ char-mouth-open.png saved")

	// Generate mouth-closed image using the open one as reference
	fmt.Println("Generating mouth-closed gopher...")
	closedContents := []*genai.Content{{
		Role: genai.RoleUser,
		Parts: []*genai.Part{
			genai.NewPartFromText("Same character, same style, same pose — but with the mouth closed in a calm, friendly smile."),
			genai.NewPartFromBytes(openData, "image/png"),
		},
	}}
	respClosed, err := client.Models.GenerateContent(ctx, model, closedContents, cfg)
	if err != nil {
		log.Fatal("mouth-closed:", err)
	}
	closedData, err := extractImage(respClosed)
	if err != nil {
		log.Fatal("mouth-closed extract:", err)
	}
	if err := os.WriteFile(outDir+"/char-mouth-closed.png", closedData, 0644); err != nil {
		log.Fatal(err)
	}
	fmt.Println("✓ char-mouth-closed.png saved")
	fmt.Println("Done! Images saved to", outDir)
}
