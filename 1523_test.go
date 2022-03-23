package lcode

import "testing"

func Test1523(t *testing.T) {
	countOdds := func(low int, high int) int {
		ret := 0
		for i := low; i <= high; i++ {
			if i%2 != 0 {
				ret++
			}
		}
		return ret
	}

	countOdds = func(low int, high int) int {
		ret := (high - low) / 2

		if low%2 != 0 && high%2 != 0 {
			ret++

			return ret
		}

		if low%2 != 0 {
			ret++
		}

		if high%2 != 0 {
			ret++
		}
		return ret
	}
	t.Log(countOdds(3, 7))
	t.Log(countOdds(8, 10))
}
