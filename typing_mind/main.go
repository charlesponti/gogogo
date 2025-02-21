package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)

type InputData struct {
	Data struct {
		Chats []Chat `json:"chats"`
	} `json:"data"`
}

type Chat struct {
	ChatID      string    `json:"chatID"`
	ChatTitle   string    `json:"chatTitle"`
	CreatedAt   time.Time `json:"createdAt"`
	Messages    []Message `json:"messages"`
	Model       string    `json:"model"`
	Preview     string    `json:"preview"`
	SyncedAt    *string   `json:"syncedAt"`
	TokenUsage  Usage     `json:"tokenUsage"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Message struct {
	// Add message fields if needed
}

type Usage struct {
	TotalTokens int `json:"totalTokens"`
}

type OutputConversation struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	Preview   string    `json:"preview"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Usage     Usage     `json:"usage"`
}

type Config struct {
	inputPath  string
	outputPath string
}

func main() {
	config := Config{}
	flag.StringVar(&config.inputPath, "input", "", "Path to the input JSON file")
	flag.StringVar(&config.outputPath, "output", "", "Path to the output JSON file")
	flag.Parse()
	
	if (config.inputPath == "" && config.outputPath == "") {
		fmt.Println("Arguments missing: input, output")
		fmt.Printf("Usage: %s -input <input_file> -output <output_file>\n", os.Args[0])
		fmt.Println("   -input   Path to the input JSON file")
		fmt.Println("   -output  Path to the output JSON file")
		os.Exit(1)
	}

	if (config.inputPath == "") {
		fmt.Println("Error: Input file path is required")
		fmt.Printf("Usage: %s -input <input_file> -output <output_file>\n", os.Args[0])
		fmt.Println("   -input   Path to the input JSON file")
		os.Exit(1)
	}

	if (config.outputPath == "") {
		fmt.Println("Error: Output file path is required")
		fmt.Printf("Usage: %s -input <input_file> -output <output_file>\n", os.Args[0])
		fmt.Println("   -input   Path to the input JSON file")
		fmt.Println("   -output  Path to the output JSON file")
		os.Exit(1)
	}


	// Read input file
	inputData, err := os.ReadFile(config.inputPath)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	// Parse input JSON
	var input InputData
	if err := json.Unmarshal(inputData, &input); err != nil {
		fmt.Printf("Error parsing input JSON: %v\n", err)
		return
	}

	// Convert to output format
	conversations := make([]OutputConversation, len(input.Data.Chats))
	for i, chat := range input.Data.Chats {
		conversations[i] = OutputConversation{
			ID:        chat.ChatID,
			Title:     chat.ChatTitle,
			Model:     chat.Model,
			Messages:  chat.Messages,
			Preview:   chat.Preview,
			CreatedAt: chat.CreatedAt,
			UpdatedAt: chat.UpdatedAt,
			Usage:     chat.TokenUsage,
		}
	}

	// Generate output JSON
	output, err := json.MarshalIndent(conversations, "", "  ")
	if err != nil {
		fmt.Printf("Error generating output JSON: %v\n", err)
		return
	}

	// Write output file
	if err := os.WriteFile(config.outputPath, output, 0644); err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		return
	}

	fmt.Println("Conversion completed successfully!")
}
