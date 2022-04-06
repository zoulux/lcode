package lcode

import (
	"math"
	"math/bits"
	"testing"
)

func Test762(t *testing.T) {
	t.Log(1 << 2 & 665772)
	t.Log(1<<2, 665772)

	t.Log("===")

	countPrimeSetBits := func(left int, right int) int {
		countBit := func(num int) int {
			return bits.OnesCount(uint(num))
		}

		mm := make(map[int]bool)
		isPrime := func(count int) (ret bool) {
			if v, ok := mm[count]; ok {
				return v
			}

			defer func() {
				mm[count] = ret
			}()

			if count == 1 {
				return false
			}

			for i := 2; i <= int(math.Sqrt(float64(count))); i++ {
				if count%i == 0 {
					return false
				}
			}

			return true
		}
		c := 0
		for i := left; i <= right; i++ {
			cb := countBit(i)
			if isPrime(cb) {
				c++
			}
		}
		return c
	}

	countPrimeSetBits = func(left int, right int) int {
		table := map[int]struct{}{
			2:  {},
			3:  {},
			5:  {},
			7:  {},
			11: {},
			13: {},
			17: {},
			19: {},
		}

		isPrime := func(count int) bool {
			if _, ok := table[count]; ok {
				return true
			}
			return false
		}

		ans := 0
		for i := left; i <= right; i++ {
			cb := bits.OnesCount(uint(i))
			if isPrime(cb) {
				ans++
			}
		}
		return ans
	}
	countPrimeSetBits = func(left int, right int) int {
		table := [...]int{
			0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 0, 1,
		}

		isPrime := func(count int) bool {
			if table[count] == 1 {
				return true
			}
			return false
		}

		ans := 0
		for i := left; i <= right; i++ {
			cb := bits.OnesCount(uint(i))
			if isPrime(cb) {
				ans++
			}
		}
		return ans
	}

	//t.Log(countPrimeSetBits(6, 10))
	//t.Log(countPrimeSetBits(10, 15))
	t.Log(countPrimeSetBits(842, 888))

}
