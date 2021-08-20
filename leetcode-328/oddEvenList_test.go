package leetcode_328

import (
	"github.com/DMwangnima/go-leetcode/model"
	"testing"
)

var (
	cases = [][]int{
		{1, 2, 3, 4},
		{1, 2, 3, 4, 5},
	}
	expects = [][]int{
		{1, 3, 2, 4},
		{1, 3, 5, 2, 4},
	}
)

func TestOddEvenList(t *testing.T) {
    for i, c := range cases {
    	list := model.NewList(c...)
    	res := oddEvenList(list)
    	t.Log(model.IsEqual(res, expects[i]...))
	}
}
