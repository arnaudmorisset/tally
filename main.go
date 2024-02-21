package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/arnaudmorisset/tally/internal/counting"
	"github.com/arnaudmorisset/tally/internal/formatting"
)

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

	output, err := formatting.FormatOutput(count, *outputFormat)
	if err != nil {
		return fmt.Errorf("Error formatting output: %v", err)
	}
	print(output)

	return nil
}

func printError(err error) {
	fmt.Fprint(os.Stderr, err.Error()+"\n")
}

func print(args ...interface{}) {
	fmt.Println(args...)
}
