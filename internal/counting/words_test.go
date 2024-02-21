package counting_test

import (
	"testing"

	"github.com/arnaudmorisset/tally/internal/counting"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		content  []byte
		expected int
	}{
		{
			name:     "Empty content",
			content:  []byte(""),
			expected: 0,
		},
		{
			name:     "Single word",
			content:  []byte("Hello"),
			expected: 1,
		},
		{
			name:     "Multiple words",
			content:  []byte("Hello, world!"),
			expected: 2,
		},
		{
			name:     "Words with whitespace",
			content:  []byte("   Hello   world   "),
			expected: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := counting.CountWords(test.content)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if result != test.expected {
				t.Errorf("Expected %d words, but got %d", test.expected, result)
			}
		})
	}
}
