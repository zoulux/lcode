package lcode

import "testing"

func Test704(t *testing.T) {

	search := func(nums []int, target int) int {
		i, j := 0, len(nums)-1

		for i <= j {
			mid := i + (j-i)>>1
			if target > nums[mid] {
				i = mid + 1
			} else if target < nums[mid] {
				j = mid - 1
			} else {
				return mid
			}
		}
		return -1
	}

	t.Log(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	t.Log(search([]int{-1, 0, 3, 5, 9, 12}, 2))

}
