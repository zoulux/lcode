package lcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if l.Next != nil {
		return fmt.Sprintf("%v -> %v", l.Val, l.Next.String())
	}
	return fmt.Sprintf("%v", l.Val)
}

func generateList(arr []int) *ListNode {
	dummy := new(ListNode)

	head := dummy
	for _, v := range arr {
		head.Next = &ListNode{Val: v}
		head = head.Next
	}

	return dummy.Next
}
