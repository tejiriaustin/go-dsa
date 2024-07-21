package two_pointer_techniques

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		name     string
		input    func() []int
		expected int
	}{
		{
			name: "successful",
			input: func() []int {
				return []int{1, 1, 2}
			},
			expected: 2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := RemoveDuplicates(testCase.input())
			if got != testCase.expected {
				t.Errorf("expected [%v], got [%v]", got, testCase.expected)
			}
		})
	}
}
