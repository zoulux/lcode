package lcode

import "testing"

func Test2206(t *testing.T) {
	divideArray := func(nums []int) bool {
		mm := make([]*struct{}, len(nums))
		ans := 0
		for i, n1 := range nums {
			if mm[i] != nil {
				continue
			}
			for j := i + 1; j < len(nums); j++ {
				if mm[j] != nil {
					continue
				}
				n2 := nums[j]
				if n1 == n2 {
					ans++
					mm[i] = &struct{}{}
					mm[j] = &struct{}{}
					break
				}
			}
		}
		return ans == len(nums)/2
	}

	divideArray = func(nums []int) bool {
		maxv := nums[0]
		for _, num := range nums {
			if num > maxv {
				maxv = num
			}
		}
		arr := make([]int, maxv)

		for _, num := range nums {
			arr[num]++
		}

		for _, v := range arr {
			if v%2 != 0 {
				return false
			}
		}
		return true
	}

	divideArray = func(nums []int) bool {
		mm := make(map[int]int)
		for _, num := range nums {
			mm[num] += 1
		}

		for _, v := range mm {
			if v%2 != 0 {
				return false
			}
		}
		return true
	}

	t.Log(divideArray([]int{3, 2, 3, 2, 2, 2}))
	t.Log(divideArray([]int{1, 2, 3, 4}))
}
