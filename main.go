package main

import (
	"fmt"
	"os"
)

// printError prints the error to stderr.
func printError(err error) {
	fmt.Fprint(os.Stderr, err.Error())
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
	print("Hello, World!")
	return nil
}
