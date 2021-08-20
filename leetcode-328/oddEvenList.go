package leetcode_328

import "github.com/DMwangnima/go-leetcode/model"

func oddEvenList(head *model.ListNode) *model.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var sli []*model.ListNode
	cur, tail := head, head
	for cur != nil && cur.Next != nil {
		sli = append(sli, cur.Next)
		if cur.Next.Next != nil {
			cur.Next = cur.Next.Next
			cur = cur.Next
			tail = cur
		} else {
			cur = cur.Next
		}
	}
	cur = tail
	for _, node := range sli {
		cur.Next = node
		cur = node
	}
	cur.Next = nil
	return head
}

// 采用一奇一偶的方式迭代，由于最后可能会有一个单独的奇节点，所以需要单独考虑
func oddEvenList1(head *model.ListNode) *model.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	oddDummy, evenDummy := &model.ListNode{Next: head}, &model.ListNode{Next: head.Next}
	oddPrev, evenPrev := oddDummy, evenDummy
	cur := head
	for cur != nil && cur.Next != nil {
		oddPrev.Next = cur
		evenPrev.Next = cur.Next
		oddPrev = cur
		evenPrev = cur.Next
		cur = cur.Next.Next
	}
	if cur != nil {
		oddPrev.Next = cur
		oddPrev = cur
		evenPrev.Next = nil
	}
	oddPrev.Next = evenDummy.Next
	return oddDummy.Next
}

// 采用一偶一奇的方式迭代，即使最后有一个单独的偶节点，也不需要单独考虑，代码更加优雅
func oddEvenList2(head *model.ListNode) *model.ListNode {
	if head == nil {
		return nil
	}
	evenHead := head.Next
	odd, even := head, evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}