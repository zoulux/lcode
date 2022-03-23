package lcode

import "testing"

func isValidSudoku(board [][]byte) bool {

	for i := range board {
		tmp1 := make([]byte, 10)
		tmp2 := make([]byte, 10)
		for j := range board[i] {
			cur1 := board[i][j]

			cur2 := board[j][i]

			if cur1 != '.' {
				if tmp1[cur1-'0'] == 0 {
					tmp1[cur1-'0'] = 1
				} else {
					return false
				}
			}

			if cur2 != '.' {
				if tmp2[cur2-'0'] == 0 {
					tmp2[cur2-'0'] = 1
				} else {
					return false
				}
			}
		}
	}

	i, j := 0, 0
	for chunk := 0; chunk < 9; chunk++ {

		tmp2 := make([]byte, 10)

		for step := 0; step < 9; step++ {
			cur := board[i][j]

			if cur != '.' {
				if tmp2[cur-'0'] == 0 {
					tmp2[cur-'0'] = 1
				} else {
					return false
				}
			}
			i++
			if i%3 == 0 {
				j++
				i = chunk * 3
			}
		}
		j = chunk
	}

	return true
}

func Test36(t *testing.T) {

}
