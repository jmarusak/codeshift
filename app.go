package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	// Retrieve API key from environment variable
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "Error: Missing GEMINI_API_KEY environment variable")
		return
	}

	// Create a GenerativeAI client with the API key
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v", err)
		return
	}

	// Get the "gemini-1.5-flash" model
	model := client.GenerativeModel("gemini-1.5-flash")

	// Generate content with the prompt
	prompt := "Why meaning of life is 42?"
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating content: %v", err)
		return
	}

	// Access the generated text
	printResponse(resp)
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("---")
}
