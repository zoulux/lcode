package lcode

import "testing"

func Test153(t *testing.T) {
	findMin := func(nums []int) int {
		left, right := 0, len(nums)-1
		for left < right {
			mid := int(uint(left+right) >> 1)
			if nums[mid] < nums[right] {
				right = mid
			} else {
				left = mid + 1
			}
		}
		return nums[left]
	}

	t.Log(findMin([]int{3, 4, 5, 1, 2}))
	t.Log(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	t.Log(findMin([]int{11, 13, 15, 17}))
}
