package leetcode_388

import "testing"

var (
	cases = []string {
		"dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext",
		"dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext",
		"a",
		"file1.txt\nfile2.txt\nlongfile.txt",
		"dir\n        file.txt",
	}
	expects = []int {
		32,
		20,
		0,
		12,
		16,
	}
)

func TestLengthLongestPath(t *testing.T) {
    for i, c := range cases {
    	res := lengthLongestPath(c)
    	if res != expects[i] {
    		t.Logf("res: %d, expect: %d", res, expects[i])
		}
	}
}
