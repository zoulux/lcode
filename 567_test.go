package lcode

import (
	"testing"
)

func Test567(t *testing.T) {
	checkInclusion := func(s1 string, s2 string) bool {
		n1, n2 := len(s1), len(s2)
		data := make(map[byte]int)

		for _, v := range []byte(s1) {
			data[v]++
		}

		win := make(map[byte]int)

		for i := 0; i < n2-n1; i++ {
			for j := i; j < i+n1; j++ {
				win[s2[j]]++
			}

			flag := true
			for k, v := range data {
				if vv, ok := win[k]; !ok || vv != v {
					flag = false
				}
			}
			if flag {
				return true
			}

			win[s2[i]]--
		}

		return false
	}

	t.Log(checkInclusion("ab", "eidbaooo"))
	t.Log(checkInclusion("ab", "eidboaoo"))
}
