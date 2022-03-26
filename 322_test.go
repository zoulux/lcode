package lcode

import (
	"math"
	"testing"
)

func Test322(t *testing.T) {
	coinChange := func(coins []int, amount int) int {
		min := func(a, b int) int {
			if a < b {
				return a
			}
			return b
		}

		get := func(arr []int, idx int) int {
			if idx >= 0 {
				return arr[idx]
			}
			return math.MaxInt
		}

		dp := make([]int, amount+1)

		for i := 1; i <= amount; i++ {
			mmin := math.MaxInt // 本轮最小值
			for _, coin := range coins {
				// 10 =>min( 10-1 ,10-2, 10-5 ) => min (9,8,5) => min(5+2+2,5+2+1,5)
				// => min(3,3,1) => 1
				mmin = min(mmin, get(dp, i-coin))
			}
			if mmin != math.MaxInt {
				// 本轮获取到最小值 则加上本轮的一枚的一枚硬币
				dp[i] = mmin + 1
			} else {
				// 本轮没有获取到，逻辑不成立
				dp[i] = math.MaxInt
			}
		}

		if dp[amount] == math.MaxInt {
			return -1
		}

		return dp[amount]
	}

	t.Log(coinChange([]int{1, 2, 5}, 100))

	if coinChange([]int{1, 2, 5}, 11) != 3 {
		t.Fatal()
	}

	if coinChange([]int{2}, 3) != -1 {
		t.Fatal()
	}

	if coinChange([]int{1, 2, 5}, 11) != 3 {
		t.Fatal()
	}

	if coinChange([]int{1, 2, 5}, 0) != 0 {
		t.Fatal()
	}

}
