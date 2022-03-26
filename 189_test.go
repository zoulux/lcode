package lcode

import (
	"testing"
)

func Test189(t *testing.T) {
	//rotate := func(nums []int, k int) {
	//	ln := len(nums)
	//	k = k % ln
	//	newNums := append(
	//		nums[ln-k:],
	//		nums[:ln-k]...,
	//	)
	//	for i := range nums {
	//		nums[i] = newNums[i]
	//	}
	//}

	reverse := func(nums []int, i, j int) {
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
	rotate := func(nums []int, k int) {
		k = k % len(nums)
		reverse(nums, 0, len(nums)-1)
		reverse(nums, 0, k-1)
		reverse(nums, k, len(nums)-1)
	}

	s1 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(s1, 3)
	t.Log(s1)
	//reverse(s1, 0, len(s1)-1)
	//t.Log(s1)

	//5 6 7 1 2 3 4

	//1, 2, 3, 4, 5, 6, 7
	// 7,6,5,4,3,2,1
	// 5,6,7  1,2,3,4

}
