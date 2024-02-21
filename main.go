package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

// printError prints the error to stderr.
func printError(err error) {
	fmt.Fprint(os.Stderr, err.Error()+"\n")
}

// print prints the arguments to stdout.
func print(args ...interface{}) {
	fmt.Println(args...)
}

type Count struct {
	Words int `json:"words"`
	Lines int `json:"lines"`
	Bytes int `json:"bytes"`
}

func main() {
	if err := run(); err != nil {
		printError(err)
		os.Exit(1)
	}
}

func run() error {
	outputFormat := flag.String("f", "text", "Output format: text, json")
	flag.Parse()

	// Read file content
	filePath := flag.Arg(0)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("Error reading file: %v", err)
	}

	// Perform counting
	count := Count{}
	count.Bytes = len(content)

	count.Words, err = countWords(content)
	if err != nil {
		return fmt.Errorf("Error counting words: %v", err)
	}

	count.Lines, err = countLines(content)
	if err != nil {
		return fmt.Errorf("Error counting lines: %v", err)
	}

	switch *outputFormat {
	case "json":
		jsonOutput, err := toJSON(count)
		if err != nil {
			return fmt.Errorf("Error converting to JSON: %v", err)
		}

		print(jsonOutput)
	case "text":
		fallthrough
	default:
		print("Words:", count.Words)
		print("Lines:", count.Lines)
		print("Bytes:", count.Bytes)
	}

	return nil
}

func countWords(content []byte) (int, error) {
	wordCount := 0

	scanner := bufio.NewScanner(bytes.NewReader(content))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return wordCount, nil
}

func countLines(content []byte) (int, error) {
	lineCount := 0

	scanner := bufio.NewScanner(bytes.NewReader(content))
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lineCount, nil
}

func toJSON(count Count) (string, error) {
	data, err := json.Marshal(count)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
