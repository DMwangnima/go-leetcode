package leetcode_5

import (
	"testing"
)

var (
	cases = []string{
		"babad",
		"cbbd",
		"ac",
		"a",
		"aacabdkacaa",
	}
	expects = []string{
		"bab",
		"bb",
		"a",
		"a",
		"aca",
	}
)

func TestLongestPalindrome(t *testing.T) {
    for i, c := range cases {
    	res := longestPalindrome(c)
    	if res != expects[i] {
    		t.Logf("res: %s, expect: %s", res, expects[i])
		}
	}
}

func TestLongestPalindrome1(t *testing.T) {
	for i, c := range cases {
		res := longestPalindrome1(c)
		t.Logf("exiting, %s", c)
		if res != expects[i] {
			t.Logf("res: %s, expect: %s", res, expects[i])
		}
	}
}
