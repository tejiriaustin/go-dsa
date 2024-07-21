package two_pointer_techniques

import (
	"reflect"
	"testing"
)

func TestMidpoint(t *testing.T) {
	testCases := []struct {
		name         string
		input        func() *LinkedList
		expected     *LinkedList
		expectsError bool
	}{
		{
			name: "success",
			input: func() *LinkedList {
				head := &LinkedList{Val: 1}
				head.Next = &LinkedList{Val: 2}
				head.Next.Next = &LinkedList{Val: 3}
				head.Next.Next.Next = &LinkedList{Val: 4}
				head.Next.Next.Next.Next = &LinkedList{Val: 5}

				return head
			},
			expected: &LinkedList{Val: 3},
		},
	}

	for _, testCase := range testCases {

		input := testCase.input()
		got := Midpoint(input)
		if !reflect.DeepEqual(testCase.expected.Val, got.Val) {
			t.Errorf("expects %v got %v", testCase.expected, got)
		}
	}
}
