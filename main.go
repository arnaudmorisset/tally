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

	if flag.NArg() != 1 {
		printHelp()
		return nil
	}

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

func printHelp() {
	fmt.Println("Usage: tally [options] <file>")
	fmt.Println("Options:")
	fmt.Println("  -f string")
	fmt.Println("        Output format: text, json")
	fmt.Println("Examples:")
	fmt.Println("  # Count words, lines, and bytes in a file (default output format is text)")
	fmt.Println("  tally filename.txt")
	fmt.Println("  # Count words, lines, and bytes in a file and export results in JSON format")
	fmt.Println("  tally -f json filename.txt")
}
