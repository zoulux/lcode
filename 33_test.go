package lcode

import (
	"testing"
)

func Test33(t *testing.T) {
	search := func(nums []int, target int) int {
		if len(nums) <= 1 {
			if nums[0] == target {
				return 0
			}
			return -1
		}

		bound := 0
		for i := 0; i+1 < len(nums); i++ {
			if nums[i] > nums[i+1] {
				bound = i + 1
			}
		}

		bsearch := func(nums []int) int {
			left, right := 0, len(nums)-1
			t := -1
			for left <= right {
				mid := int(uint(left+right) >> 1) // avoid overflow when computing h
				if target < nums[mid] {
					right = mid - 1
				} else if target > nums[mid] {
					left = mid + 1
				} else {
					return mid
				}
			}
			return t
		}

		x := bsearch(nums[:bound])
		if x != -1 {
			return x
		}
		x = bsearch(nums[bound:])
		if x != -1 {
			return x + bound
		}
		return -1
	}
	search = func(nums []int, target int) int {
		left, right := 0, len(nums)-1
		for left <= right {
			mid := int(uint(left+right) >> 1) // avoid overflow when computing h
			if target == nums[mid] {
				return mid
			} else if target == nums[left] {
				return left
			} else if target == nums[right] {
				return right
			} else if nums[left] < nums[mid] {
				if target > nums[left] && target < nums[mid] {
					right = mid - 1
				} else {
					left = mid + 1
				}
			} else {
				if target > nums[mid] && target < nums[right] {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}

		}
		return -1
	}

	t.Log(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	//t.Log(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	//t.Log(search([]int{1}, 0))
}
