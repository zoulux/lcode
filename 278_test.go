package lcode

import "testing"

func Test278(t *testing.T) {
	isBadVersion := func(version int) bool { return false }

	firstBadVersion := func(n int) int {
		i, j := 1, n
		for i <= j {
			mid := i + (j-i)>>1
			if mv := isBadVersion(mid); mv {
				j = mid - 1
			} else {
				i = mid + 1
			}
		}
		return i
	}

	t.Log(firstBadVersion(5))

}
