package lcode

import "testing"

func Test70(t *testing.T) {
	climbStairs := func(n int) int {
		dp := make([]int, n+1)

		if n >= 1 {
			dp[1] = 1
		}
		if n >= 2 {
			dp[2] = 2
		}
		for i := 3; i <= n; i++ {
			dp[i] = dp[i-1] + dp[i-2]
		}

		return dp[n]
	}

	t.Log(climbStairs(1))
	t.Log(climbStairs(2))
	t.Log(climbStairs(3))
	t.Log(climbStairs(45))
}
