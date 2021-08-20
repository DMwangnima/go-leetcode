package leetcode_237

import "github.com/DMwangnima/go-leetcode/model"

// 一开始想到的方法
func deleteNode(node *model.ListNode) {
	prev, cur := node, node.Next
	for cur.Next != nil {
		prev.Val = cur.Val
		prev = cur
		cur = cur.Next
	}
	prev.Val = cur.Val
	prev.Next = nil
}

// 其实很明显，这道题的前提条件在于要删除的节点不能是最后一个节点
func deleteNode1(node *model.ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
