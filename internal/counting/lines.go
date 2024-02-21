package counting

import (
	"bufio"
	"bytes"
)

// CountLines counts the number of lines in the content.
func CountLines(content []byte) (int, error) {
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
