package lcode

import (
	"testing"
)

func twoSum(nums []int, target int) []int {
	nl := len(nums)
	for i := 0; i < nl; i++ {
		for j := i + 1; j < nl; j++ {
			ni, nj := nums[i], nums[j]
			if ni+nj == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for idx, elem := range nums {
		if v, ok := m[target-elem]; ok {
			return []int{v, idx}
		}
		m[elem] = idx
	}
	return nil
}

func Test1(t *testing.T) {
	//t.Log(twoSum([]int{2, 2, 11, 7, 2, 11, 15}, 17))
	//t.Log(twoSum([]int{3, 2, 4}, 6))
	t.Log(twoSum([]int{3, 2, 4}, 6))
	var arr []int
	for i := 0; i <= 10000; i++ {
		arr = append(arr, i)
	}

	t.Log(twoSum(arr, 19999))

}
