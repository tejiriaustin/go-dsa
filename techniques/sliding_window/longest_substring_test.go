package sliding_window

import "testing"

func TestLongestSubstring(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "success",
			input:    "abcabcbb",
			expected: 3,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := longestSubstring(test.input)
			if got != test.expected {
				t.Errorf("expected %v got %v", test.expected, got)
			}
		})
	}
}
