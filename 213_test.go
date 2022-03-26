package lcode

import (
	"testing"
)

func Test213(t *testing.T) {

	rob := func(nums []int) int {
		if len(nums) == 0 {
			return 0
		}
		if len(nums) == 1 {
			return nums[0]
		}

		max := func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}

		get := func(dp []int, idx int) int {
			if idx >= 0 {
				return dp[idx]
			}
			return 0
		}

		_rob := func(nums []int) int {
			dp := make([]int, len(nums)) // 经过 i 房间时是获得最大的收益
			ret := nums[0]
			for i := 0; i < len(nums); i++ {
				dp[i] = max(get(dp, i-2)+nums[i], get(dp, i-1))
				ret = max(ret, dp[i])
			}
			return ret
		}

		return max(_rob(nums[:len(nums)-1]), _rob(nums[1:]))
	}
	t.Log(rob([]int{2})) // 2
	//t.Log(rob([]int{1, 2, 3, 1})) // 4
	//t.Log(rob([]int{1, 2, 3}))    // 4

	//t.Log(rob([]int{1, 2, 3, 1}))          // 4
	//t.Log(rob([]int{2, 7, 1, 3, 20})) // 27
	//t.Log(rob([]int{2, 7, 9, 3, 1}))       // 12
	//t.Log(rob([]int{1}))                   // 1
	//t.Log(rob([]int{1, 2}))                // 2
	//t.Log(rob([]int{1, 2, 3}))             // 4
	//t.Log(rob([]int{4, 1, 2, 7, 5, 3, 1})) // 14
}
