package two_pointer_techniques

import (
	"reflect"
	"testing"
)

func TestMoveZero(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "success",
			input:    []int{1, 0, 2, 0, 4, 6},
			expected: []int{1, 2, 4, 6, 0, 0},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := moveZero(testCase.input)
			if !reflect.DeepEqual(got, testCase.expected) {
				t.Errorf("expected %v got %v", testCase.expected, got)
			}
		})
	}
}
