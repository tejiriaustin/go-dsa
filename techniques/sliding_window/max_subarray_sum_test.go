package sliding_window

import "testing"

func TestMaxSubArraySum(t *testing.T) {
	testCases := []struct {
		name        string
		input       []int
		k           int
		expectedSum int
	}{
		{
			name:        "success",
			input:       []int{1, 3, 2, 5, 1, 1, 6, 2},
			k:           3,
			expectedSum: 12,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := maxSubArraySum(test.input, test.k)
			if got != test.expectedSum {
				t.Errorf("expected %v got %v", test.expectedSum, got)
			}
		})
	}
}
