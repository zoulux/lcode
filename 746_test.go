package lcode

import "testing"

func Test746(t *testing.T) {

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	minCostClimbingStairs := func(cost []int) int {
		dp := make([]int, len(cost)+1)
		dp[0] = cost[0]
		dp[1] = cost[1]

		for i := 2; i < len(cost); i++ {
			dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
		}

		return min(dp[len(cost)-1], dp[len(cost)-2])
	}

	t.Log(minCostClimbingStairs([]int{10, 15, 20}))
	t.Log(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}
