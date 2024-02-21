package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/arnaudmorisset/tally/internal/counting"
	"github.com/arnaudmorisset/tally/internal/formatting"
)

// printError prints the error to stderr.
func printError(err error) {
	fmt.Fprint(os.Stderr, err.Error()+"\n")
}

// print prints the arguments to stdout.
func print(args ...interface{}) {
	fmt.Println(args...)
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

	count := formatting.Count{}
	count.Bytes = len(content)

	count.Words, err = counting.CountWords(content)
	if err != nil {
		return fmt.Errorf("Error counting words: %v", err)
	}

	count.Lines, err = counting.CountLines(content)
	if err != nil {
		return fmt.Errorf("Error counting lines: %v", err)
	}

	switch *outputFormat {
	case "json":
		jsonOutput, err := formatting.ToJSON(count)
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
