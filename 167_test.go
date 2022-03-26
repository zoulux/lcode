package lcode

import "testing"

func Test167(t *testing.T) {
	twoSum := func(numbers []int, target int) []int {
		for i := 0; i < len(numbers); i++ {
			sub := target - numbers[i]
			for j := i + 1; j < len(numbers); j++ {
				if numbers[j] == sub {
					return []int{i + 1, j + 1}
				}
			}
		}
		return nil
	}

	twoSum = func(numbers []int, target int) []int {
		for i := 0; i < len(numbers); i++ {
			sub := target - numbers[i]
			for left, right := i+1, len(numbers)-1; left <= right; {
				mid := (right-left)>>1 + left
				if sub > numbers[mid] {
					left = mid + 1
				} else if sub < numbers[mid] {
					right = mid - 1
				} else {
					return []int{i + 1, mid + 1}
				}
			}
		}
		return nil
	}

	twoSum = func(numbers []int, target int) []int {
		for left, right := 0, len(numbers)-1; left <= right; {
			sum := numbers[left] + numbers[right]
			if sum > target {
				right--
			} else if sum < target {
				left++
			} else {
				return []int{left + 1, right + 1}
			}
		}
		return nil
	}

	t.Log(twoSum([]int{2, 7, 11, 15}, 9))
	t.Log(twoSum([]int{2, 3, 4}, 6))
	t.Log(twoSum([]int{-1, 0}, -1))
}
