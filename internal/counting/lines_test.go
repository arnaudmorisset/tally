package counting_test

import (
	"testing"

	"github.com/arnaudmorisset/tally/internal/counting"
)

func TestCountLines(t *testing.T) {
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
			name:     "Single line",
			content:  []byte("Hello, world!"),
			expected: 1,
		},
		{
			name:     "Multiple lines",
			content:  []byte("Line 1\nLine 2\nLine 3"),
			expected: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := counting.CountLines(test.content)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if result != test.expected {
				t.Errorf("Expected %d lines, but got %d", test.expected, result)
			}
		})
	}
}
