package leetcode_160

import "github.com/DMwangnima/go-leetcode/model"

// 直接使用哈希表，时间复杂度O(m+n)，空间复杂度O(m)
func getIntersectionNode(headA, headB *model.ListNode) *model.ListNode {
	m := make(map[*model.ListNode]struct{})
	for headA != nil {
		m[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := m[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

// 看官方题解
func getIntersectionNode1(headA, headB *model.ListNode) *model.ListNode {
    if headA == nil || headB == nil {
    	return nil
	}
	curA, curB := headA, headB
	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}

		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA
}
