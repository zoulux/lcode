package lcode

import "testing"

func Test19(t *testing.T) {
	removeNthFromEnd := func(head *ListNode, n int) *ListNode {
		length, dummy := 0, head
		for dummy != nil {
			length++
			dummy = dummy.Next
		}

		remove := func(head *ListNode, n int) *ListNode {
			dummy := new(ListNode)
			dummy.Next = head

			h := dummy
			for i := 0; i < length-n; i++ {
				h = h.Next
			}

			if h.Next != nil {
				h.Next = h.Next.Next
			}

			return dummy.Next
		}

		return remove(head, n)
	}
	removeNthFromEnd = func(head *ListNode, n int) *ListNode {
		dummy := &ListNode{Next: head}
		slow, fast := dummy, head

		// 快指针先走 n 步
		for i := 0; i < n; i++ {
			fast = fast.Next
		}

		// 慢指针跟着快指针走
		// 快指针到结尾， 慢指针敲好指向倒数第 n 个节点
		for ; fast != nil; fast = fast.Next {
			slow = slow.Next
		}
		slow.Next = slow.Next.Next
		return dummy.Next
	}

	t.Log(removeNthFromEnd(generateList([]int{1, 2, 3, 4, 5}), 2)) // 1 -> 2 -> 3 -> 5
	t.Log(removeNthFromEnd(generateList([]int{1, 2, 3, 4, 5}), 1)) //1 -> 2 -> 3 -> 4
	t.Log(removeNthFromEnd(generateList([]int{1}), 1))             // nil
	t.Log(removeNthFromEnd(generateList([]int{1, 2}), 1))          //1
	t.Log(removeNthFromEnd(generateList([]int{1, 2}), 2))          //2
}
