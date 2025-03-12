package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type PullSecret struct {
	Auths map[string]interface{} `json:"auths"`
}

func main() {
	// Define command line flags
	inputDir := flag.String("dir", "", "Directory containing JSON pull secret files")
	outputFile := flag.String("output", "combined-pull-secret.json", "Output file path")
	flag.Parse()

	if *inputDir == "" {
		fmt.Println("Error: Please specify input directory using -dir flag")
		flag.Usage()
		os.Exit(1)
	}

	// Create combined pull secret structure
	combinedSecret := PullSecret{
		Auths: make(map[string]interface{}),
	}

	// Read all JSON files from the input directory
	files, err := os.ReadDir(*inputDir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		os.Exit(1)
	}

	// Process each JSON file
	for _, file := range files {
		if filepath.Ext(file.Name()) != ".json" {
			continue
		}

		filePath := filepath.Join(*inputDir, file.Name())
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", file.Name(), err)
			continue
		}

		var pullSecret PullSecret
		if err := json.Unmarshal(content, &pullSecret); err != nil {
			fmt.Printf("Error parsing JSON from file %s: %v\n", file.Name(), err)
			continue
		}

		// Merge auths into combined secret
		for registry, auth := range pullSecret.Auths {
			combinedSecret.Auths[registry] = auth
		}
	}

	// Convert combined secret to JSON
	outputJSON, err := json.MarshalIndent(combinedSecret, "", "  ")
	if err != nil {
		fmt.Printf("Error creating JSON output: %v\n", err)
		os.Exit(1)
	}

	// Write to output file
	if err := os.WriteFile(*outputFile, outputJSON, 0644); err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully combined pull secrets into %s\n", *outputFile)
}
