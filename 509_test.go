package lcode

import "testing"

func Test509(t *testing.T) {
	fib := func(n int) int {
		a, b := 0, 1
		for i := 0; i < n; i++ {
			a, b = b, a+b
		}
		return a
	}

	fib = func(n int) int {
		dp := make([]int, n+1)
		if n >= 1 {
			dp[1] = 1
		}

		for i := 2; i <= n; i++ {
			dp[i] = dp[i-1] + dp[i-2]
		}
		return dp[n]
	}

	t.Log(fib(0))
	t.Log(fib(1))
	t.Log(fib(2))
	t.Log(fib(3))
	t.Log(fib(4))
}
