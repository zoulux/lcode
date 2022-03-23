package lcode

import (
	"sort"
	"testing"
)

func Test1491(t *testing.T) {
	average := func(salary []int) float64 {
		sort.Ints(salary)
		salary = append(salary[1:])
		salary = append(salary[:len(salary)-1])

		var ret float64 = 0
		for _, s := range salary {
			ret += float64(s)
		}

		return ret / float64(len(salary))
	}
	average = func(salary []int) float64 {
		max, min := 0, 1<<63-1
		ret := 0
		for _, s := range salary {
			if s < min {
				min = s
			}

			if s > max {
				max = s
			}
			ret += s
		}

		ret -= max
		ret -= min

		return float64(ret) / float64(len(salary)-2)
	}
	t.Log(average([]int{4000, 3000, 1000, 2000}))
	t.Log(average([]int{1000, 2000, 3000}))
	t.Log(average([]int{6000, 5000, 4000, 3000, 2000, 1000}))
	t.Log(average([]int{8000, 9000, 2000, 3000, 6000, 1000}))
	t.Log(average([]int{1, 2, 3, 4}))
}
