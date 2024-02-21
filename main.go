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
		err := fmt.Errorf("Error reading file: %v", err)
		printError(err)
		os.Exit(1)
	}

	// Perform counting
	count := Count{
		Words: countWords(content),
		Lines: countLines(content),
		Bytes: len(content),
	}

	switch *outputFormat {
	case "json":
		jsonOutput, _ := toJSON(count)
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

func countWords(content []byte) int {
	wordCount := 0

	scanner := bufio.NewScanner(bytes.NewReader(content))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordCount++
	}

	if err := scanner.Err(); err != nil {
		printError(err)
	}

	return wordCount
}

func countLines(content []byte) int {
	lineCount := 0

	scanner := bufio.NewScanner(bytes.NewReader(content))
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		printError(err)
	}

	return lineCount
}

func toJSON(count Count) (string, error) {
	data, err := json.Marshal(count)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
