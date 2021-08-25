package leetcode_316

import "testing"

var (
	cases = []string {
		"bcabc",
		"adcad",
		"cbacdcbc",
		"bbcdd",
	}
	expects = []string {
		"abc",
		"acd",
		"acdb",
		"bcd",
	}
)

func TestRemoveDuplicateLetters(t *testing.T) {
    for i, c := range cases {
    	res := removeDuplicateLetters(c)
    	if res != expects[i] {
    		t.Logf("res: %s, expect: %s", res, expects[i])
		}
	}
}

func TestRemoveDuplicateLetters1(t *testing.T) {
	for i, c := range cases {
		res := removeDuplicateLetters1(c)
		if res != expects[i] {
			t.Logf("res: %s, expect: %s", res, expects[i])
		}
	}
}
