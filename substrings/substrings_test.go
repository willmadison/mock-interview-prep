package substrings_test

import (
	"testing"
	"fmt"
	"github.com/willmadison/mock-interview-prep/substrings"
)

func TestLongestUnique(t *testing.T) {
	cases := []struct{
		given string
		expected string
	}{
		{
			given: "mom",
			expected: "mo",
		},
		{
			given: "modsy",
			expected: "modsy",
		},
		{
			given: "longwall",
			expected: "longwa",
		},
		{
			given: "d",
			expected: "d",
		},
		{
			given: "evaryonen",
			expected: "evaryon",
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("substrings.LongestUnique(%s)", tc.given), func (t *testing.T) {
			actual := substrings.LongestUnique(tc.given)

			if actual != tc.expected {
				t.Errorf("got: %s; wanted: %s", actual, tc.expected)
			}
		})
	}

}