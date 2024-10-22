package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/v66/github"
	"golang.org/x/oauth2"
)

// Recursive function to list all files in a repository directory
func listFilesInRepo(client *github.Client, owner, repo, path, branch string) error {
	// Get the contents of the directory or file
	fileContent, directoryContent, _, err := client.Repositories.GetContents(
		context.Background(),
		owner, repo, path,
		&github.RepositoryContentGetOptions{
			Ref: branch,
		},
	)
	if err != nil {
		return err
	}

	// If contents is not nil, it's a file
	if fileContent != nil {
		// It's a file, print the file content
		content, err := fileContent.GetContent()
		if err != nil {
			return err
		}
		fmt.Printf("File in: %s\n", *fileContent.Path)
		fmt.Print(content)
		return nil
	}

	// Otherwise, we are dealing with a directory, so iterate through directoryContent
	for _, content := range directoryContent {
		if *content.Type == "file" {
			// If it's a code file, ask for the file content
			if strings.HasSuffix(*content.Path, ".js") {
				err := listFilesInRepo(client, owner, repo, *content.Path, branch)
				if err != nil {
					return err
				}
			}
		} else if *content.Type == "dir" {
			// If it's a directory, recurse into it
			err := listFilesInRepo(client, owner, repo, *content.Path, branch)
			if err != nil {
				return err
			}
		}
	}

	return nil
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

	// Define the owner, repository, root directory, and branch
	owner := "pepinoloco"
	repo := "devops"
	root := ""      // Root directory, an empty string means the root of the repo
	branch := "main" // or any other branch

	// List all files in the repository
	err := listFilesInRepo(client, owner, repo, root, branch)
	if err != nil {
		log.Fatalf("Error listing files: %v", err)
	}
}

