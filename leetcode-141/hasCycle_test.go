package leetcode_141

import (
	"github.com/DMwangnima/go-leetcode/model"
	"testing"
)

var (
	cases = [][]int {
	    {-1},
	    {1, -1}, // has not cycle
	    {1, 0}, // has cycle
	    {1, 2, 3, 4, 5, 0}, // cycle index 0
	    {1, 2, 3, 4, 5, 3}, // cycle index 3
    }
    expects = []bool {
    	false,
    	false,
    	true,
    	true,
    	true,
	}
)

func TestHasCycle(t *testing.T) {
    for i, c := range cases {
    	list := model.NewCycleList(c[len(c)-1], c[:len(c)-1]...)
    	if hasCycle(list) != expects[i] {
    		t.Log(c)
    		t.Log(expects[i])
		}
	}
}

func TestHasCycle1(t *testing.T) {
	for i, c := range cases {
		list := model.NewCycleList(c[len(c)-1], c[:len(c)-1]...)
		if hasCycle1(list) != expects[i] {
			t.Log(c)
			t.Log(expects[i])
		}
	}
}
