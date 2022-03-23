package lcode

import "testing"

func Test1137(t *testing.T) {
	tribonacci := func(n int) int {
		dp := make([]int, n+1)
		if n >= 1 {
			dp[1] = 1
		}
		if n >= 2 {
			dp[2] = 1
		}

		for i := 3; i <= n; i++ {
			dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
		}
		return dp[n]
	}

	t.Log(tribonacci(0))
	t.Log(tribonacci(1))
	t.Log(tribonacci(2))
	t.Log(tribonacci(3))
	t.Log(tribonacci(4))
	t.Log(tribonacci(25))

}
