package lcode

import "testing"

func Test162(t *testing.T) {
	findPeakElement := func(nums []int) int {

		left, right := 0, len(nums)-1
		for left < right {
			mid := int(uint(left+right) >> 1)
			if nums[mid] < nums[mid+1] { // 下坡不一定遇到峰值
				left = mid + 1
			} else {
				right = mid
			}
		}
		return left
	}

	t.Log(findPeakElement([]int{1, 2, 3, 1}))
	t.Log(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}
