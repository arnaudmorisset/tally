package formatting

import (
	"encoding/json"
	"fmt"
)

// Count represents the count of words, lines and bytes.
type Count struct {
	Words int `json:"words"`
	Lines int `json:"lines"`
	Bytes int `json:"bytes"`
}

// FormatOutput returns the count structure as a string with the given format.
func FormatOutput(count Count, format string) (string, error) {
	switch format {
	case "json":
		return toJSON(count)
	case "text":
		fallthrough
	default:
		return toText(count), nil
	}
}

func toText(count Count) string {
	return fmt.Sprintf("Words: %d\nLines: %d\nBytes: %d\n", count.Words, count.Lines, count.Bytes)
}

func toJSON(count Count) (string, error) {
	data, err := json.Marshal(count)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
