package lcode

import (
	"math"
	"testing"
)

func Test6071(t *testing.T) {
	minimumRounds := func(tasks []int) int {
		dp := func(n int) int {
			dp := make([]int, n+1)

			if n >= 1 {
				dp[1] = math.MaxInt
			}
			if n >= 2 {
				dp[2] = 1
			}

			if n >= 3 {
				dp[3] = 1
			}

			for i := 4; i <= n; i++ {
				a2 := dp[i-2]
				a3 := dp[i-3]

				min := a2

				if a3 < a2 {
					min = a3
				}

				dp[i] = min + 1
			}

			return dp[n]
		}

		m := make(map[int]int)
		for _, t := range tasks {
			m[t]++
		}

		count := 0
		for _, v := range m {
			if dp(v) == math.MaxInt {
				return -1
			}
			count += dp(v)
		}
		return count
	}

	t.Log(minimumRounds([]int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}))
	t.Log(minimumRounds([]int{2, 3, 3}))
}
