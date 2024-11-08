package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/DIMO-Network/model-garage/pkg/ruptela"
)

// MessageWrapper represents the structure of each message in the JSON file
type MessageWrapper struct {
	Messages []Message `json:"messages"`
}

// Message represents the structure of individual messages
type Message struct {
	Format    string          `json:"format"`
	Topic     string          `json:"topic"`
	Timestamp int64           `json:"timestamp"`
	Payload   json.RawMessage `json:"payload"`
	QOS       int             `json:"qos"`
}

func main() {
	// Define command-line flags
	inputFile := flag.String("input", "", "Path to input JSON file")
	outFilePath := flag.String("output", "-", "Path to output JSON file")
	flag.Parse()

	if *inputFile == "" {
		fmt.Println("Please provide an input file path using -input flag")
		os.Exit(1)
	}
	outFile := os.Stdout
	if *outFilePath != "-" {
		var err error
		outFile, err = os.Create(*outFilePath)
		if err != nil {
			fmt.Printf("Error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer outFile.Close()
	}

	// Open and read the input file
	file, err := os.Open(*inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Read the entire file
	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Parse the wrapper structure
	var wrapper MessageWrapper
	err = json.Unmarshal(bytes, &wrapper)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		os.Exit(1)
	}
	if len(wrapper.Messages) == 0 {
		wrapper.Messages = []Message{
			{
				Payload: bytes,
			},
		}
	}

	// Process each message
	for i, msg := range wrapper.Messages {
		// Process the message using DecodeStatusSignals
		data, errs := ruptela.ReplaceSignalsFromV1Data(msg.Payload)
		if len(errs) != 0 {
			log.Fatalf("Error decoding message %d: %v\n", i+1, errs)
		}
		wrapper.Messages[i].Payload = data
	}
	//marshal and write the updated JSON to the output file
	updatedJSON, err := json.MarshalIndent(wrapper, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling updated JSON: %v\n", err)
		os.Exit(1)
	}
	_, err = outFile.Write(updatedJSON)
	if err != nil {
		fmt.Printf("Error writing updated JSON: %v\n", err)
		os.Exit(1)
	}
}
