package lcode

import (
	"strings"
	"testing"
)

func Test557(t *testing.T) {
	reverseWords := func(s string) string {
		reverseString := func(s []byte) {
			for i, j := 0, len(s)-1; i < j; {
				s[i], s[j] = s[j], s[i]
				i++
				j--
			}
		}

		arr := strings.Split(s, " ")
		for i, v := range arr {
			tmp := []byte(v)
			reverseString(tmp)
			arr[i] = string(tmp)
		}
		return strings.Join(arr, " ")
	}
	reverseWords = func(s string) string {
		sb := []byte(s)

		reverseString := func(s []byte) {
			for i, j := 0, len(s)-1; i < j; {
				s[i], s[j] = s[j], s[i]
				i++
				j--
			}
		}

		for i, j := 0, 0; i < len(sb); i = j {
			for j < len(sb) && sb[j] != ' ' {
				j++
			}
			reverseString(sb[i:j])
			j++ // 跳到下一个单词的头
		}
		return string(sb)
	}

	t.Log(reverseWords("Let's take LeetCode contest"))
	t.Log(reverseWords("God Ding"))
}
