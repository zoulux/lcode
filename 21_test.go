package lcode

import "testing"

func Test21(t *testing.T) {
	mergeTwoLists := func(list1 *ListNode, list2 *ListNode) *ListNode {
		dummy := new(ListNode)
		head := dummy
		for list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				head.Next = list1
				list1 = list1.Next
			} else {
				head.Next = list2
				list2 = list2.Next
			}
			head = head.Next
		}
		if list1 != nil {
			head.Next = list1
		}
		if list2 != nil {
			head.Next = list2
		}
		return dummy.Next
	}
	mergeTwoLists = func(list1 *ListNode, list2 *ListNode) *ListNode {
		if list1 == nil {
			return list2
		}
		if list2 == nil {
			return list1
		}

		if list1.Val > list2.Val {
			return mergeTwoLists(list2, list1)
		}

		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}

	t.Log(mergeTwoLists(generateList([]int{1, 2, 4}), generateList([]int{1, 3, 4})))

}
