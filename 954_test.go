package lcode

import (
	"sort"
	"testing"
)

func Test954(t *testing.T) {
	//给定一个长度为偶数的整数数组 arr，只有对 arr 进行重组后可以满足 “对于每个 0 <=i < len(arr) / 2，
	//都有 arr[2 * i + 1] = 2 * arr[2 * i]”时，返回 true；否则，返回 false。

	canReorderDoubled := func(arr []int) bool {
		sort.Ints(arr)

		flag := true
		for i := 0; i < len(arr)/2; i += 2 {
			if arr[2*i+1] != 2*arr[2*i] {
				flag = false
			}
		}

		return flag
	}

	t.Log(canReorderDoubled([]int{3, 1, 3, 6}))   //false
	t.Log(canReorderDoubled([]int{2, 1, 2, 6}))   //false
	t.Log(canReorderDoubled([]int{4, -2, 2, -4})) //true
	t.Log(canReorderDoubled([]int{-2, -4, 2, 4})) //true
}
