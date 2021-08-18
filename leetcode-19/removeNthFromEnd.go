package leetcode_19

import "github.com/DMwangnima/go-leetcode/model"

// 最直接方式，先遍历拿到总数量m，再找到m-n这个节点，进行删除
func removeNthFromEnd(head *model.ListNode, n int) *model.ListNode {
	cur := head
	sum := 0
	for cur != nil {
		sum += 1
		cur = cur.Next
	}
	dummy := &model.ListNode{Next: head}
	count := 0
	prev := dummy
	index := sum - n
	cur = head
	for cur != nil && count < index {
		prev = cur
		cur = cur.Next
		count += 1
	}
	prev.Next = prev.Next.Next
	return dummy.Next
}

// 双指针法
func removeNthFromEnd1(head *model.ListNode, n int) *model.ListNode {
    dummy := &model.ListNode{Next: head}
    front, back := head, dummy
    for i := 0; i < n; i++ {
    	front = front.Next
	}
	for front != nil {
		front = front.Next
		back = back.Next
	}
	back.Next = back.Next.Next
	return dummy.Next
}

// 使用栈
func removeNthFromEnd2(head *model.ListNode, n int) *model.ListNode {
	var sta []*model.ListNode
	dummy := &model.ListNode{Next: head}
	head = dummy
	for head != nil {
		sta = append(sta, head)
		head = head.Next
	}
	prev := sta[len(sta)-n-1]
	prev.Next = prev.Next.Next
	return dummy.Next
}
