package counting

import (
	"bufio"
	"bytes"
)

// CountWords counts the number of words in the content.
func CountWords(content []byte) (int, error) {
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
