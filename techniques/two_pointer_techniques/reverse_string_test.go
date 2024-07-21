package two_pointer_techniques

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "successful",
			input:    "qwertyuiop",
			expected: "poiuytrewq",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := ReverseString(testCase.input)
			if got != testCase.expected {
				t.Errorf("expected %s, got %s", testCase.expected, got)
			}
		})
	}
}
