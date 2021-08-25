package leetcode_71

import "testing"

var (
	cases = []string{
		"/home/",
		"/",
		"/a/b/./../c/",
		"/../",
	}
	expects = []string{
		"/home",
		"/",
		"/a/c",
		"/",
	}
)

func TestSimplifyPath(t *testing.T) {
    for i, c := range cases {
    	res := simplifyPath(c)
    	if res != expects[i] {
    		t.Logf("res: %s, expects: %s", res, expects[i])
		}
	}
}

func TestSimplifyPath1(t *testing.T) {
	for i, c := range cases {
		res := simplifyPath1(c)
		if res != expects[i] {
			t.Logf("res: %s, expects: %s", res, expects[i])
		}
	}
}