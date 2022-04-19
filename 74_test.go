package lcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test74(t *testing.T) {
	searchMatrix := func(matrix [][]int, target int) bool {
		col0 := make([]int, len(matrix))
		for i, row := range matrix {
			col0[i] = row[0]
		}

		x := sort.SearchInts(col0, target)
		if x == len(col0) {
			x--
		} else if col0[x] == target {
			return true
		} else if x > 0 {
			x--
		}

		y := sort.SearchInts(matrix[x], target)
		if y == len(matrix[x]) || matrix[x][y] != target {
			return false
		}
		return true
	}

	assert.Equal(t, searchMatrix([][]int{
		{},
	},
		3), false)

	assert.Equal(t, searchMatrix([][]int{
		{1},
	},
		3), false)

	assert.Equal(t, searchMatrix([][]int{
		{1, 3},
	},
		3), true)

	assert.Equal(t, searchMatrix([][]int{
		{1},
	}, 0), false)

	assert.Equal(t, searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, 1), true)

	assert.Equal(t, searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, 2), false)

	assert.Equal(t, searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, 3), true)

	assert.Equal(t, searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, 7), true)

	assert.Equal(t, searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, 9), false)

	assert.Equal(t, searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, 61), false)

	assert.Equal(t, searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, 35), false)

}
