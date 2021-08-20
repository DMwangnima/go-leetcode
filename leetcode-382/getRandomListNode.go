package leetcode_382

import (
	"github.com/DMwangnima/go-leetcode/model"
	"math/rand"
)

type Solution struct {
	sli []*model.ListNode
}

func Constructor(head *model.ListNode) Solution {
	s := Solution{}
	for head != nil {
		s.sli = append(s.sli, head)
		head = head.Next
	}
	return s
}

func (s *Solution) GetRandom() int {
	return s.sli[rand.Intn(len(s.sli))].Val
}

type Solution1 struct {
	head *model.ListNode
	length int
}

func Constructor1(head *model.ListNode) Solution1 {
	s := Solution1{head: head}
	for cur := head; cur != nil; cur = cur.Next {
		s.length += 1
	}
	return s
}

func (s *Solution1) GetRandom() int {
	step := rand.Intn(s.length) + 1
	res := s.head
	for i := 1; i < step; i++ {
		res = res.Next
	}
	return res.Val
}

// 蓄水池抽样算法，需要再学习一下
type Solution2 struct {

}
