package formatting_test

import (
	"testing"

	"github.com/arnaudmorisset/tally/internal/formatting"
)

func TestFormatOutput(t *testing.T) {
	tests := []struct {
		name        string
		count       formatting.Count
		format      string
		expected    string
		expectedErr error
	}{
		{
			name:        "JSON format",
			count:       formatting.Count{Words: 10, Lines: 5, Bytes: 100},
			format:      "json",
			expected:    `{"words":10,"lines":5,"bytes":100}`,
			expectedErr: nil,
		},
		{
			name:        "Text format",
			count:       formatting.Count{Words: 10, Lines: 5, Bytes: 100},
			format:      "text",
			expected:    "Words: 10\nLines: 5\nBytes: 100\n",
			expectedErr: nil,
		},
		{
			name:        "Default format",
			count:       formatting.Count{Words: 10, Lines: 5, Bytes: 100},
			format:      "unknown",
			expected:    "Words: 10\nLines: 5\nBytes: 100\n",
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := formatting.FormatOutput(test.count, test.format)
			if err != test.expectedErr {
				t.Errorf("Unexpected error. Expected: %v, Got: %v", test.expectedErr, err)
			}
			if result != test.expected {
				t.Errorf("Unexpected result. Expected: %s, Got: %s", test.expected, result)
			}
		})
	}
}
