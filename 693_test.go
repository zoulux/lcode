package lcode

import (
	"testing"
)

func Test693(t *testing.T) {
	hasAlternatingBits := func(n int) bool {
		n1 := (n >> 1) | n
		return (n>>1)&n == 0 && (n1&(n1+1) == 0)
	}
	t.Log(hasAlternatingBits(5))  //true
	t.Log(hasAlternatingBits(7))  //false
	t.Log(hasAlternatingBits(11)) //false
	t.Log(hasAlternatingBits(4))  // false

	t.Logf("%b , %b ,%b", 0b11, 0b11+1, 0b11&(0b11+1))
}
