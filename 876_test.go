package lcode

import "testing"

func Test876(t *testing.T) {
	middleNode := func(head *ListNode) *ListNode {
		slow, fast := head, head
		for fast != nil && fast.Next != nil {
			fast = fast.Next
			slow = slow.Next

			if fast != nil {
				fast = fast.Next
			}
		}
		return slow
	}

	//t.Log(middleNode(generateList([]int{1, 2, 3, 4, 5})))
	t.Log(middleNode(generateList([]int{1, 2, 3, 4})))

}
