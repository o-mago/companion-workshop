package main

import (
	"context"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/model/gemini"
)

const instruction = `You are bot, a friendly, helpful AI companion.
Answer no more than 3 sentences.`

func newRootAgent(ctx context.Context) (agent.Agent, error) {
	m, err := gemini.NewModel(ctx, "gemini-2.0-flash", nil)
	if err != nil {
		return nil, err
	}

	return llmagent.New(llmagent.Config{
		Name:        "companion_agent",
		Description: "A friendly and engaging AI companion.",
		Instruction: instruction,
		Model:       m,
	})
}
