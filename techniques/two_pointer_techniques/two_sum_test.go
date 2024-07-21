package two_pointer_techniques

import (
	"reflect"
	"testing"
)

func TestTwoSums(t *testing.T) {
	testCases := []struct {
		name  string
		input struct {
			arr    []int
			target int
		}
		expected func() []int
	}{
		{
			name: "successful",
			input: struct {
				arr    []int
				target int
			}{arr: []int{2, 7, 11, 15}, target: 9},
			expected: func() []int {
				return []int{0, 1}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := twoSum(testCase.input.arr, testCase.input.target)
			expected := testCase.expected()
			if !reflect.DeepEqual(got, expected) {
				t.Errorf("expected %v, got %v", expected, got)
			}
		})
	}
}
