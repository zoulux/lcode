package lcode

import "testing"

func Test2210(t *testing.T) {
	countHillValley := func(nums []int) int {
		if len(nums) <= 2 {
			return 0
		}

		ans := 0
		for i := 1; i < len(nums)-1; i++ {
			if nums[i] == nums[i-1] {
				continue
			}

			j := i + 1
			for j < len(nums) {
				if nums[i] != nums[j] {
					break
				}
				j++
			}

			if nums[i] > nums[i-1] {
				if j < len(nums) && nums[i] > nums[j] {
					ans++
				}
			} else {
				if j < len(nums) && nums[i] < nums[j] {
					ans++
				}
			}
		}
		return ans
	}
	t.Log(countHillValley([]int{2, 4, 1, 1, 6, 5, 5})) //3
	//t.Log(countHillValley([]int{2, 4, 1, 1, 6, 5}))                                                                                               //3
	//t.Log(countHillValley([]int{6, 6, 5, 5, 4, 1}))                                                                                               //0
	t.Log(countHillValley([]int{57, 57, 57, 57, 57, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 90, 85, 85, 85, 86, 86, 86})) //0

}
