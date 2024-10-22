package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v66/github"
	"golang.org/x/oauth2"
)

// Fetches file content from a GitHub repository
func getFileContent(client *github.Client, owner string, repo string, path string, branch string) ([]byte, error) {
	fileContent, _, _, err := client.Repositories.GetContents(
		context.Background(), owner, repo, path, &github.RepositoryContentGetOptions{
			Ref: branch, // optional, can specify branch or commit
		},
	)
	if err != nil {
		return nil, err
	}

	// Get the content in raw format (decoded)
	content, err := fileContent.GetContent()
	if err != nil {
		return nil, err
	}

	return []byte(content), nil
}

func main() {
	// GitHub access token (set it as an environment variable for security)
	token := os.Getenv("GITHUB_TOKEN")

	// Set up the OAuth2 token and HTTP client
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// Define the owner, repository, file path, and branch
	owner := "pepinoloco"
	repo := "devops"
	path := "README.md"
	branch := "main" // or any other branch

	// Fetch the file content
	content, err := getFileContent(client, owner, repo, path, branch)
	if err != nil {
		log.Fatalf("Error fetching file: %v", err)
	}

	fmt.Printf("File content:\n%s\n", content)
}

