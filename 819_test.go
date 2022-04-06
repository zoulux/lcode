package lcode

import (
	"strings"
	"testing"
)

func Test819(t *testing.T) {
	mostCommonWord := func(paragraph string, banned []string) string {

		sign := []string{"!", "?", "'", ",", ";", "."}

		trim := func(word string) string {
			for _, s := range sign {
				word = strings.Trim(word, s)
			}
			word = strings.ToLower(word)
			return word
		}
		words := strings.FieldsFunc(paragraph, func(r rune) bool {
			for _, s := range sign {
				if r == rune(s[0]) {
					return true
				}
			}
			if r == ' ' {
				return true
			}
			return false

		})

		bannedMap := make(map[string]struct{})
		for _, word := range banned {
			word = trim(word)
			bannedMap[word] = struct{}{}
		}

		m := make(map[string]int)
		for _, word := range words {
			word = trim(word)
			if _, ok := bannedMap[word]; ok {
				continue
			}
			m[word]++
		}

		max := 0
		maxs := ""
		for k, v := range m {
			if v > max {
				max = v
				maxs = k
			}
		}
		return maxs
	}
	t.Log(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
	t.Log(mostCommonWord("a, a, a, a, b,b,b,c, c", []string{"a"}))
}
