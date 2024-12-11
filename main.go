package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"bytes"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// API request struct
type PromptRequest struct {
	Message string `json:"message"`
}

// Google Gemini API client and model instance
var (
	ctx    context.Context
	client *genai.Client
	model  *genai.GenerativeModel
)

// Middleware to handle CORS
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight OPTIONS request
		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Read the body of the request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Parse the JSON body into a Go struct
	var promptRequest PromptRequest
	err = json.Unmarshal(body, &promptRequest)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Actuall prompt in the request
	prompt := promptRequest.Message

	// Call the GenerativeAI API to generate content
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating content: %v", err)
		return
	}

	// Extract the response from the API response
	var response bytes.Buffer
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				if txt, ok := part.(genai.Text); ok {
					response.WriteString(string(txt))
				}
			}
		}
	}

	fmt.Println(strings.Repeat("*", 100))
	fmt.Print(prompt)
	fmt.Println(strings.Repeat("-", 100))
	fmt.Print(response.String())

	// Write the response back to the client
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response.Bytes())
	if err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func initGemini() error {
	ctx = context.Background()

	// Retrieve API key from environment variable
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("missing GEMINI_API_KEY environment variable")
	}

	// Create a GenerativeAI client with the API key
	var err error
	client, err = genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return fmt.Errorf("error creating client: %v", err)
	}

	// Get the "gemini-1.5-flash" model
	model = client.GenerativeModel("gemini-1.5-flash")

	// Set response MIME type to JSON
	model.ResponseMIMEType = "application/json"
	return nil
}

func main() {
	mux := http.NewServeMux()
	handler := corsMiddleware(mux)

	// Serve static files from the "./static" directory at the "/static" path
	staticDir := http.FileServer(http.Dir("./static"))
	mux.Handle("/", staticDir)

	mux.HandleFunc("/generate", generateHandler)

	// Initialize context and client only once
	if err := initGemini(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing Gemini client: %v", err)
		return
	}

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
