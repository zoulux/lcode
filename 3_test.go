package lcode

import "testing"

func Test3(t *testing.T) {
	lengthOfLongestSubstring := func(s string) int {
		n := len(s)
		left, right := 0, 0
		res := 0
		win := make(map[byte]int)
		for right < n {
			c := s[right]
			right++
			win[c]++

			for win[c] > 1 {
				d := s[left]
				left++
				win[d]--
			}

			if diff := right - left; diff > res {
				res = diff
			}
		}

		return res
	}

	t.Log(lengthOfLongestSubstring("abcabcbb")) //3
	t.Log(lengthOfLongestSubstring("bbbbb"))    //1

}
