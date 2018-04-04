package palindrome_test

import (
	"testing"
	"github.com/willmadison/mock-interview-prep/palindrome"
	"fmt"
)

func TestIsPalindrome(t *testing.T) {
	cases := []struct{
		given string
		expected bool
	}{
		{
			given: "mom",
			expected: true,
		},
		{
			given: "mommy",
		},
		{
			given: "this is a long word",
		},
		{
			given: "d",
			expected: true,
		},
		{
			given: "",
			expected: true,
		},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("palindrome.IsA(%s)", tc.given), func (t *testing.T) {
			actual := palindrome.IsA(tc.given)

			if actual != tc.expected {
				t.Errorf("got: %v; wanted: %v", actual, tc.expected)
			}
		})
	}

}
