package leetcode_86

import "github.com/DMwangnima/go-leetcode/model"

// 自己的思路，有点像快速排序的思路。先找到链表的尾节点，之后以此遍历，如果大于等于x，则插入到最后
func partition(head *model.ListNode, x int) *model.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &model.ListNode{Next: head}
	var preTail *model.ListNode
	cur := dummy
	for cur.Next != nil {
		cur = cur.Next
	}
	preTail = cur
	cur = dummy.Next
	prev, tail := dummy, preTail
	for cur != preTail {
		if cur.Val >= x {
			prev.Next = cur.Next
			tail.Next = cur
			cur.Next = nil
			tail = tail.Next
			cur = prev.Next
		} else {
			prev = prev.Next
			cur = cur.Next
		}
	}
	// 为保证相对顺序，尾节点可能也需要插入到最后
	if cur.Val >= x && cur.Next != nil {
		prev.Next = cur.Next
		tail.Next = cur
		cur.Next = nil
	}
	return dummy.Next
}

func partition1(head *model.ListNode, x int) *model.ListNode {
	smallHead, bigHead := &model.ListNode{Next: head}, &model.ListNode{Next: head}
	small, big := smallHead, bigHead
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			big.Next = head
			big = big.Next
		}
		head = head.Next
	}
	// 这句一定要在后一句的前面 假如链表为[1] x=2，顺序倒转就会出错
	big.Next = nil
	small.Next = bigHead.Next
	return smallHead.Next
}
