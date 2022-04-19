package lcode

import (
	"sort"
	"strconv"
	"testing"
)

func Test386(t *testing.T) {
	lexicalOrder := func(n int) []int {
		arr := make([]int, 0, n+1)
		for i := 1; i <= n; i++ {
			arr = append(arr, i)
		}

		sort.SliceStable(arr, func(i, j int) bool {
			vi := strconv.Itoa(arr[i])
			vj := strconv.Itoa(arr[j])
			return vi < vj
		})
		return arr
	}

	lexicalOrder = func(n int) []int {
		arr := make([]int, 0, n+1)
		var dfs func(int)
		dfs = func(cur int) {
			if cur > n {
				return
			}
			arr = append(arr, cur)
			for i := 0; i <= 9; i++ {
				dfs(cur*10 + i)
			}
		}

		for i := 1; i <= 9; i++ {
			dfs(i)
		}
		return arr
	}

	t.Log(lexicalOrder(13)) // [1,10,11,12,13,2,3,4,5,6,7,8,9]
	t.Log(lexicalOrder(2))
	t.Log(lexicalOrder(23))
	t.Log(lexicalOrder(100)) //
	t.Log(lexicalOrder(134))
}
