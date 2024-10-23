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

/*
	// Define the schema
	schema := `{
    "content": [
        {
            "type": "text",
            "value": "string"
        },
        {
            "type": "code",
            "language": "string",
            "value": "string"
        }
    ]
	}`

*/

	// Generate content with the prompt
	prompt := "generate simple main.go module for golang"


	// Ask the model to respond with JSON.
	model.ResponseMIMEType = "application/json"
	// Specify the schema.
	model.ResponseSchema = &genai.Schema{
        Type:  genai.TypeArray,
        Items: &genai.Schema{Type: genai.TypeString},
	}

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
