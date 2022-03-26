package lcode

import (
	"testing"
)

func Test740(t *testing.T) {
	deleteAndEarn := func(nums []int) int {
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

		maxNum := nums[0]

		for _, num := range nums {
			if num > maxNum {
				maxNum = num
			}
		}

		all := make([]int, maxNum+1)
		for _, num := range nums {
			all[num]++
		}

		dp := make([]int, len(all)) // 经过第 i 个数时获得的最大点数

		ret := nums[0]
		for i := 0; i < len(all); i++ {
			dp[i] = max(
				get(dp, i-2)+all[i]*i, // 选择当前值
				get(dp, i-1),          // 不选择当前值
			)
			ret = max(ret, dp[i])
		}

		return ret
	}
	deleteAndEarn = func(nums []int) int {
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

		maxNum := nums[0]

		for _, num := range nums {
			if num > maxNum {
				maxNum = num
			}
		}

		dp := make([]int, maxNum+1)
		for _, num := range nums {
			dp[num] += num
		}

		ret := nums[0]
		for i := 0; i < len(dp); i++ {
			dp[i] = max(
				get(dp, i-2)+dp[i], // 选择当前值
				get(dp, i-1),       // 不选择当前值
			)
			ret = max(ret, dp[i])
		}

		return ret
	}
	t.Log(deleteAndEarn([]int{3, 4, 2}))          // 6
	t.Log(deleteAndEarn([]int{2, 2, 3, 3, 3, 4})) // 9
}
