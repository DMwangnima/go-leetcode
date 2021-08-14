package model

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func NewList(vals ...int) *ListNode {
	dummy := &ListNode{}
	prev := dummy
	for _, val := range vals {
		newNode := &ListNode{Val: val}
		prev.Next = newNode
		prev = newNode
	}
	if dummy.Next == nil {
		return nil
	}
	return dummy.Next
}

func NewCycleList(index int, vals ...int) *ListNode {
	var indexNode *ListNode
	dummy := &ListNode{}
	prev := dummy
	for i, val := range vals {
		newNode := &ListNode{Val: val}
		prev.Next = newNode
		prev = newNode
		if i == index {
			indexNode = newNode
		}
	}
	prev.Next = indexNode
	return dummy.Next
}

func PrintList(head *ListNode) {
	switch {
	case head == nil:
		fmt.Println("empty list")
		return
	case head.Next == nil:
		fmt.Printf("%d\n", head.Val)
		return
	}

    for head.Next != nil {
    	fmt.Printf("%d->", head.Val)
    	head = head.Next
	}
	fmt.Println(head.Val)
}

func IsEqual(head *ListNode, vals ...int) bool {
	if head == nil {
		if len(vals) == 0 {
			return true
		}
		return false
	}
	for _, val := range vals {
		if head.Val != val {
			return false
		}
		head = head.Next
	}
	return true
}
