package leetcode_148

import (
	"github.com/DMwangnima/go-leetcode/model"
	"testing"
)

func TestSortList(t *testing.T) {
	cases := [][]int {
		{},
		{2},
		{4, 2, 1, 3},
		{5, 4, 2, 1, 3},
		{4, 4, 2, 1, 3},
	}
	expects := [][]int {
		{},
		{2},
		{1, 2, 3, 4},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 4},
	}
	for i, c := range cases {
		list := model.NewList(c...)
		newList := sortList(list)
		model.PrintList(newList)
		t.Log(expects[i])
		t.Log(model.IsEqual(newList, expects[i]...))
	}
}
