package lcode

import "testing"

func Test344(t *testing.T) {
	reverseString := func(s []byte) {
		for i, j := 0, len(s)-1; i < j; {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}
	}

	s1 := []byte("hello")
	reverseString(s1)
	t.Log(string(s1))

	s2 := []byte("hello0")
	reverseString(s2)
	t.Log(string(s2))
}
