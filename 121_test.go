package lcode

import (
	"math"
	"testing"
)

func Test121(t *testing.T) {
	maxProfit := func(prices []int) int {
		n := len(prices)
		max := 0
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if prices[j] > prices[i] {
					target := prices[j] - prices[i]
					if target > max {
						max = target
					}
				}
			}
		}
		return max
	}

	maxProfit = func(prices []int) int {
		minPrice := math.MaxInt
		mmax := 0
		for _, price := range prices {
			if price < minPrice {
				minPrice = price
			}
			if target := price - minPrice; target > mmax {
				mmax = target
			}
		}
		return mmax
	}
	maxProfit = func(prices []int) int {

		min := func(a, b int) int {
			if a < b {
				return a
			}
			return b
		}

		max := func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}

		dp := make([]int, len(prices)) // 在第 i 天得到的最低价
		dp[0] = prices[0]

		for i := 1; i < len(prices); i++ {
			dp[i] = min(prices[i], dp[i-1])
		}

		maxx := 0
		for i, price := range prices {
			maxx = max(maxx, price-dp[i])
		}
		return maxx
	}

	t.Log(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	t.Log(maxProfit([]int{7, 6, 4, 3, 1}))
	t.Log(maxProfit([]int{7, 6, 2, 3, 100}))
	t.Log(maxProfit([]int{2, 1, 2, 0, 1}))
	t.Log(maxProfit([]int{2, 1, 2, 0, 10}))
	t.Log(maxProfit([]int{2, 1, 99, 0, 10}))
}
