package leetcode_138

import "github.com/DMwangnima/go-leetcode/model"

// 使用哈希表维持原链表与新链表节点的映射关系
func copyRandomList(head *model.RandomListNode) *model.RandomListNode {
	m := make(map[*model.RandomListNode]*model.RandomListNode)
	dummy := &model.RandomListNode{}
	prev := dummy
	cur := head
	for cur != nil {
		newNode := &model.RandomListNode{Val: cur.Val}
		m[cur] = newNode
		prev.Next = newNode
		prev = newNode
		cur = cur.Next
	}
	cur = head
	curNew := dummy.Next
	for curNew != nil {
		curNew.Random = m[cur.Random]
		cur = cur.Next
		curNew = curNew.Next
	}
	return dummy.Next
}

// 既然哈希表能维持映射关系，那么使用旧节点的Next指针指向新节点，新节点的Next指针指向旧节点的下一个节点也可以维持映射关系
// 即A->B->C变为A->A'->B->B'->C->C'
func copyRandomList1(head *model.RandomListNode) *model.RandomListNode {
	if head == nil {
		return nil
	}
	for cur := head; cur != nil; cur = cur.Next.Next {
		cur.Next = &model.RandomListNode{Next: cur.Next, Val: cur.Val}
	}
	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
	}
	res := head.Next
	for cur := head; cur != nil; {
		tmp := cur.Next.Next
		if tmp != nil {
			cur.Next.Next = tmp.Next
		}
		cur.Next = tmp
		cur = tmp
	}
	return res
}
