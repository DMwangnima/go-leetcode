package leetcode_143

import (
	"github.com/DMwangnima/go-leetcode/model"
	"testing"
)

func TestReorderList(t *testing.T) {
	cases := [][]int {
		{1, 2, 3, 4},
		{1, 2, 3, 4, 5},
		{},
		{1},
	}
	expects := [][]int {
		{1, 4, 2, 3},
		{1, 5, 2, 4, 3},
		{},
		{1},
	}
	for i, c := range cases {
		list := model.NewList(c...)
		reorderList(list)
		if !model.IsEqual(list, expects[i]...) {
			model.PrintList(list)
		}
	}
}

func TestReorderList1(t *testing.T) {
	cases := [][]int {
		{1, 2, 3, 4},
		{1, 2, 3, 4, 5},
		{},
		{1},
	}
	expects := [][]int {
		{1, 4, 2, 3},
		{1, 5, 2, 4, 3},
		{},
		{1},
	}
	for i, c := range cases {
		list := model.NewList(c...)
		reorderList1(list)
		if !model.IsEqual(list, expects[i]...) {
			model.PrintList(list)
		}
	}
}
