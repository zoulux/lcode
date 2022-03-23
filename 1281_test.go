package lcode

import "testing"

func Test1281(t *testing.T) {
	subtractProductAndSum := func(n int) int {
		arr := make([]int, 0)

		for ; n != 0; n /= 10 {
			arr = append(arr, n%10)
		}

		sum, mut := 0, 1
		for _, v := range arr {
			sum += v
			mut *= v
		}
		return mut - sum
	}
	subtractProductAndSum = func(n int) int {
		sum, mut := 0, 1
		for ; n != 0; n /= 10 {
			sum += n % 10
			mut *= n % 10
		}
		return mut - sum
	}
	t.Log(subtractProductAndSum(234))
	t.Log(subtractProductAndSum(4421))
}
