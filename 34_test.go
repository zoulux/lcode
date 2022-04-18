package lcode

import (
	"testing"
)

func Test34(t *testing.T) {
	searchRange := func(nums []int, target int) []int {
		leftSearch := func(left, right int) int {
			t := -1
			for left <= right {
				mid := int(uint(left+right) >> 1) // avoid overflow when computing h
				if target < nums[mid] {
					right = mid - 1
				} else if target > nums[mid] {
					left = mid + 1
				} else {
					right, t = mid-1, mid
				}
			}
			return t
		}

		rightSearch := func(left, right int) int {
			t := -1
			for left <= right {
				mid := int(uint(left+right) >> 1) // avoid overflow when computing h
				if target < nums[mid] {
					right = mid - 1
				} else if target > nums[mid] {
					left = mid + 1
				} else {
					left, t = mid+1, mid
				}
			}
			return t
		}

		l := leftSearch(0, len(nums)-1)
		r := rightSearch(0, len(nums)-1)
		return []int{l, r}
	}

	t.Log(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	t.Log(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	t.Log(searchRange([]int{}, 6))
	t.Log(searchRange([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1))
	t.Log(searchRange([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1))
	t.Log(searchRange([]int{1, 2, 3, 3, 3, 3, 4, 5, 9}, 3))

}
