package lcode

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test821(t *testing.T) {
	shortestToChar := func(s string, c byte) (ans []int) {
		var idx []int
		for i, v := range s {
			if v == rune(c) {
				idx = append(idx, i)
			}
		}

		abs := func(a, b int) int {
			det := a - b
			if det > 0 {
				return det
			}
			return -det
		}

		min := func(a, b int) int {
			if a > b {
				return b
			}
			return a
		}

		for i := range s {
			det := math.MaxInt
			for _, v := range idx {
				det = min(det, abs(i, v))
			}

			ans = append(ans, det)
		}
		return
	}
	shortestToChar = func(s string, c byte) []int {
		min := func(a, b int) int {
			if a > b {
				return b
			}
			return a
		}

		n := len(s)
		ans := make([]int, n)

		idx := -n
		for i, ch := range s {
			if byte(ch) == c {
				idx = i
			}
			ans[i] = i - idx
		}

		idx = n * 2
		for i := n - 1; i >= 0; i-- {
			if s[i] == c {
				idx = i
			}
			ans[i] = min(ans[i], idx-i)
		}
		return ans

	}
	assert.Equal(t, shortestToChar("loveleetcode", 'e'), []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0})
}
