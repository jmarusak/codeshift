package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// File struct represents a file
type File struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// Folder struct represents a folder and contains a slice of children (which can be files or folders)
type Folder struct {
	Name     string     `json:"name"`
	Type     string     `json:"type"`
	Children []FileOrFolder `json:"children"`
}

// FileOrFolder is an interface that can represent both File and Folder
type FileOrFolder interface{}


// CreatSampleCodebase creates a sample Folder object with hardcoded data
func CreatSampleCodebase() Folder {
	// Create sample files
	file1 := File{Name: "report.txt", Type: "file"}
	file2 := File{Name: "presentation.pptx", Type: "file"}
	file3 := File{Name: "image.png", Type: "file"}
	file4 := File{Name: "notes.docx", Type: "file"}

	// Create sample folders
	subfolder1 := Folder{
		Name: "Documents",
		Type: "folder",
		Children: []FileOrFolder{
			file1,
			file2,
		},
	}

	subfolder2 := Folder{
		Name: "Images",
		Type: "folder",
		Children: []FileOrFolder{
			file3,
			file4,
		},
	}

	// Create the root folder with children
	rootFolder := Folder{
		Name: "My Files",
		Type: "folder",
		Children: []FileOrFolder{
			subfolder1,
			subfolder2,
		},
	}

	return rootFolder
}

func MarshalCodebaseToJSON(folder Folder) (string, error) {
	jsonData, err := json.MarshalIndent(folder, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func main() {
	// Create a sample folder object
	rootFolder := CreatSampleCodebase()

	// Print the resulting folder structure
	fmt.Printf("Created Folder: %+v\n", rootFolder)

	// Marshal the folder object to JSON
	jsonString, err := MarshalCodebaseToJSON(rootFolder)
	if err != nil {
		log.Fatalf("Error marshalling folder to JSON: %v", err)
	}

	// Print the resulting JSON string
	fmt.Println("JSON Output:")
	fmt.Println(jsonString)
}
