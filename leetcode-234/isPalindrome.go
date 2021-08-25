package leetcode_234

import (
	"github.com/DMwangnima/go-leetcode/model"
)

// 先使用快慢指针找到中间节点，再反转链表，最后使用双指针检验
// 这里的关键在于快慢指针的起始位置以及终止的条件
func isPalindrome(head *model.ListNode) bool {
    if head == nil {
    	return true
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	tmp := slow.Next
	slow.Next = nil
	right := reverseList(tmp)
	for right != nil {
		if head.Val != right.Val {
			return false
		}
		head = head.Next
		right = right.Next
	}
	return true
}

func reverseList(head *model.ListNode) *model.ListNode {
	var prev *model.ListNode
	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}
	return prev
}
