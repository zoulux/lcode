package lcode

import (
	"fmt"
	"testing"
)

func Test191(t *testing.T) {
	hammingWeight := func(num uint32) int {
		ret := 0
		for ; num != 0; num = num >> 1 {
			ret += int(num % 2)
		}
		return ret
	}

	hammingWeight = func(num uint32) int {
		ret := 0

		fmt.Printf("%d, %b %b %b\n", num, num, num-1, num&(num-1))
		for ; num > 0; num &= num - 1 {
			ret++
		}
		return ret
	}

	t.Log(hammingWeight(00000000000000000000000000001011))
	t.Log(hammingWeight(00000000000000000000000010000000))
	//t.Log(hammingWeight(11111111111111111111111111111101))
}
