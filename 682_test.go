package lcode

import (
	"strconv"
	"testing"
)

func Test682(t *testing.T) {
	calPoints := func(ops []string) int {
		points := make([]int, 0, len(ops))
		for _, op := range ops {
			switch op {
			case "C":
				points = points[:len(points)-1]
			case "D":
				points = append(points, points[len(points)-1]*2)
			case "+":
				points = append(points, points[len(points)-1]+points[len(points)-2])
			default:
				point, _ := strconv.Atoi(op)
				points = append(points, point)
			}
		}
		sum := 0
		for _, point := range points {
			sum += point
		}
		return sum
	}
	t.Log(calPoints([]string{"5", "2", "C", "D", "+"}))
	t.Log(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}))
	t.Log(calPoints([]string{"1"}))
}
