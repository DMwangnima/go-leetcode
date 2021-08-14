package leetcode_141

import "github.com/DMwangnima/go-leetcode/model"

// 使用哈希表
func hasCycle(head *model.ListNode) bool {
	seen := make(map[*model.ListNode]struct{})
	for head != nil {
		if _, ok := seen[head]; ok {
			return true
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return false
}

// 使用快慢指针
func hasCycle1(head *model.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head.Next, head
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}
