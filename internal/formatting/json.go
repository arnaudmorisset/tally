package formatting

import "encoding/json"

// Count represents the count of words, lines and bytes.
type Count struct {
	Words int `json:"words"`
	Lines int `json:"lines"`
	Bytes int `json:"bytes"`
}

// ToJSON converts the count struct to a JSON string.
func ToJSON(count Count) (string, error) {
	data, err := json.Marshal(count)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
