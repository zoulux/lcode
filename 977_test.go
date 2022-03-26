package lcode

import (
	"sort"
	"testing"
)

func Test977(t *testing.T) {
	sortedSquares := func(nums []int) []int {
		ret := make([]int, len(nums))
		for idx, num := range nums {
			ret[idx] = num * num
		}
		sort.Ints(ret)
		return ret
	}
	sortedSquares = func(nums []int) []int {
		i, j, k := 0, len(nums)-1, len(nums)-1

		ret := make([]int, len(nums))
		for i <= j {
			if nums[i]*nums[i] < nums[j]*nums[j] {
				ret[k] = nums[j] * nums[j]
				j--
			} else {
				ret[k] = nums[i] * nums[i]
				i++
			}
			k--
		}
		return ret
	}

	t.Log(sortedSquares([]int{-4, -1, 0, 3, 10}))
	t.Log(sortedSquares([]int{-7, -3, 2, 3, 11}))
}
