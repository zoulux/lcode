package lcode

import "testing"

func Test35(t *testing.T) {
	searchInsert := func(nums []int, target int) int {
		left, right := 0, len(nums)-1
		for left <= right {
			mid := (right-left)>>1 + left
			if target > nums[mid] {
				left = mid + 1
			} else if target < nums[mid] {
				right = mid - 1
			} else {
				return mid
			}
		}
		return left
	}

	t.Log(searchInsert([]int{1, 3, 5, 6}, 5))
	t.Log(searchInsert([]int{1, 3, 5, 6}, 2))
	t.Log(searchInsert([]int{1, 3, 5, 6}, 7))
	t.Log(searchInsert([]int{1, 3, 5, 6}, 0))
	t.Log(searchInsert([]int{1}, 0))
}
