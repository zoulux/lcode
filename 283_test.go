package lcode

import "testing"

func Test283(t *testing.T) {
	moveZeroes := func(nums []int) {
		for i, j := 0, 0; i < len(nums); i++ {
			if nums[i] != 0 {
				if i != j {
					nums[j], nums[i] = nums[i], nums[j]
				}
				j++
			}
		}
	}

	s1 := []int{0, 1, 0, 3, 12}
	moveZeroes(s1)
	t.Log(s1)

	s2 := []int{0, 0, 1}
	moveZeroes(s2)
	t.Log(s2)
}
