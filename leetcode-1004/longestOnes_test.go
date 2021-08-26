package leetcode_1004

import "testing"

var (
	cases = [][]int{
        {1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
        {0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
	}
	ks = []int{
		2,
		3,
	}
	expects = []int{
		6,
		10,
	}
)

func TestLongestOnes(t *testing.T) {
	for i, c := range cases {
		res := longestOnes(c, ks[i])
		if res != expects[i] {
			t.Logf("res: %d, expect: %d", res, expects[i])
		}
	}
}
