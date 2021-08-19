package leetcode_2

import "github.com/DMwangnima/go-leetcode/model"

// 除了最后的进位要生成新节点，其余全部复用
func addTwoNumbers(l1 *model.ListNode, l2 *model.ListNode) *model.ListNode {
	var carry bool
	var sum int
	var prev *model.ListNode
	dummy := &model.ListNode{Next: l1}
	prev = dummy
	for l1 != nil && l2 != nil {
        sum = l1.Val + l2.Val
        if carry {
        	sum += 1
		}
        if sum >= 10 {
        	l1.Val = sum - 10
        	carry = true
		} else {
			l1.Val = sum
			carry = false
		}
		prev = l1
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 != nil {
		for carry && l1 != nil{
			tmp := l1.Val + 1
			if tmp == 10 {
				l1.Val = 0
			} else {
				l1.Val = tmp
				carry = false
			}
			prev = l1
			l1 = l1.Next
		}
		if l1 == nil && carry {
			prev.Next = &model.ListNode{Val: 1}
			carry = false
		}
	}
	if l2 != nil {
		prev.Next = l2
		for carry && l2 != nil {
			tmp := l2.Val + 1
			if tmp == 10 {
				l2.Val = 0
			} else {
				l2.Val = tmp
				carry = false
			}
			prev = l2
			l2 = l2.Next
		}
		if l2 == nil && carry {
			prev.Next = &model.ListNode{Val: 1}
			carry = false
		}
	}
	if carry {
		prev.Next = &model.ListNode{Val: 1}
	}
	return dummy.Next
}

// 相较于前一种方法，抽取出公共逻辑
func addTwoNumbers1(l1 *model.ListNode, l2 *model.ListNode) *model.ListNode {
	var carry bool
	var sum int
	var prev *model.ListNode
	dummy := &model.ListNode{Next: l1}
	prev = dummy
	for l1 != nil && l2 != nil {
		sum = l1.Val + l2.Val
		if carry {
			sum += 1
		}
		if sum >= 10 {
			l1.Val = sum - 10
			carry = true
		} else {
			l1.Val = sum
			carry = false
		}
		prev = l1
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 != nil {
		concatenate(prev, l1, carry)
	}
	if l2 != nil {
		prev.Next = l2
		concatenate(prev, l2, carry)
	}
	if carry {
		prev.Next = &model.ListNode{Val: 1}
	}
	return dummy.Next
}

func concatenate(prev, head *model.ListNode, carry bool) {
	for carry && head != nil {
		tmp := head.Val + 1
		if tmp == 10 {
			head.Val = 0
		} else {
			head.Val = tmp
			carry = false
		}
		prev = head
		head = head.Next
	}
	if head == nil && carry {
		prev.Next = &model.ListNode{Val: 1}
		carry = false
	}
}

// 生成新节点
func addTwoNumbers2(l1 *model.ListNode, l2 *model.ListNode) *model.ListNode {
	dummy := &model.ListNode{}
	prev := dummy
	carry := 0
	for l1 != nil || l2 != nil {
		tmp1, tmp2 := 0, 0
		if l1 != nil {
			tmp1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp2 = l2.Val
			l2 = l2.Next
		}
		sum := tmp1 + tmp2 + carry
		carry = sum / 10
		prev.Next = &model.ListNode{Val: sum % 10}
		prev = prev.Next
	}
	if carry > 0 {
		prev.Next = &model.ListNode{Val: 1}
	}
	return dummy.Next
}

// 一年前的写法，几乎双百评价。。到底发生了什么
// 这里的核心在于计算进位和余位，不像第一种方法使用if，而是直接使用求余和求商
func addTwoNumbers3(l1 *model.ListNode, l2 *model.ListNode) *model.ListNode {
	var result = l1
	var carry int
	var frontPtr1 *model.ListNode
	for l1 != nil && l2 != nil {
		tmpSum := l1.Val + l2.Val + carry
		remain := tmpSum % 10
		carry = tmpSum / 10
		l1.Val = remain
		frontPtr1 = l1
		l1 = l1.Next
		l2 = l2.Next
	}
	if l2 != nil {
		frontPtr1.Next = l2
		l1 = l2
	}
	for carry != 0 && l1 != nil {
		tmpSum := l1.Val + carry
		remain := tmpSum % 10
		carry = tmpSum / 10
		l1.Val = remain
		frontPtr1 = l1
		l1 = l1.Next
	}
	if carry != 0 {
		newNode := &model.ListNode{
			Val:  carry,
			Next: nil,
		}
		frontPtr1.Next = newNode
	}
	return result
}