package main

import "testing"

func Test_Sum(t *testing.T) {
	// creating test cases is useful to loop through the possible scenarios
	// you want to test without having to do the setup all over again.
	testCases := []struct {
		name   string
		input  []int
		output int
	}{
		{
			name:   "testing with 1 + 3",
			input:  []int{1, 3},
			output: 4,
		},
		{
			name:   "testing with 4 + 4",
			input:  []int{4, 4},
			output: 8,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// we use the elipsis here to pass individual array elements
			result := Sum(tc.input[0], tc.input[1])
			if result != tc.output {
				t.Errorf("expected %d but got %d", tc.output, result)
			}
		})
	}
}
