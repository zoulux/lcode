package lcode

import (
	"testing"
)

func Test55(t *testing.T) {
	canJump := func(nums []int) bool {
		max := func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}

		mmax := 0
		n := len(nums)

		for i := 0; i < n; i++ {
			if i > mmax {
				return false
			}
			mmax = max(mmax, i+nums[i])

			if mmax > n-1 {
				return true
			}
		}
		return true
	}
	canJump = func(nums []int) bool {
		max := func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}

		if len(nums) == 1 {
			return true
		}

		dp := make([]int, len(nums)) //在 i 点上能跳最远的索引是多少
		dp[0] = nums[0]

		for i := 1; i < len(nums); i++ {
			// 如果上个节点不能达到这里就 false
			if dp[i-1] < i {
				return false
			}
			// 当前节点能到达最远的距离是，上个节点能到达最远的距离和当前节点到达最远距离的最大值
			dp[i] = max(dp[i-1], i+nums[i])

			// 当前节点如果能到达最后一个节点
			if dp[i] >= len(nums)-1 {
				return true
			}
		}

		return false
	}

	if canJump([]int{2, 3, 1, 1, 4}) != true {
		t.Fail()
	}

	if canJump([]int{3, 2, 1, 0, 4}) != false {
		t.Fail()
	}

	if canJump([]int{3, 2, 2, 0, 4}) != true {
		t.Fail()

	}
	if canJump([]int{0, 2, 3}) != false {
		t.Fail()
	}
}
