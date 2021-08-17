package leetcode_143

import "github.com/DMwangnima/go-leetcode/model"

// 1->2->3->4 => 1->4->2->3
// 1->2->3->4->5 => 1->5->2->4->3
func reorderList(head *model.ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	fast, slow := head.Next, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	tmp := slow.Next
	slow.Next = nil
	rightHead := reverseList(tmp)
	leftHead := head
	for leftHead != nil {
		tmpLeft := leftHead.Next
		tmpRight := rightHead.Next
		if tmpLeft != nil {
			rightHead.Next = tmpLeft
		}
		leftHead.Next = rightHead
		leftHead = tmpLeft
		rightHead = tmpRight
	}
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

// 先转化为切片，重点关注一下双指针移动的边界条件
func reorderList1(head *model.ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	var arr []*model.ListNode
	cur := head
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.Next
	}
	i, j := 0, len(arr)-1
	for i < j {
		arr[i].Next = arr[j]
		i += 1
		if i == j {
			break
		}
		arr[j].Next = arr[i]
		j -= 1
	}
}
