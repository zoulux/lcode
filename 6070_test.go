package lcode

import (
	"strconv"
	"strings"
	"testing"
)

func Test6070(t *testing.T) {
	digitSum := func(s string, k int) string {
		var _digitSum func(s string, k int) string
		_digitSum = func(s string, k int) string {
			if len(s) <= k {
				return s
			}
			var sb strings.Builder
			for i := 0; i < len(s); i += k {
				sum := 0
				for j := i; j < len(s) && j < i+k; j++ {
					sum += int(s[j] - '0')
				}
				sb.WriteString(strconv.Itoa(sum))
			}

			return _digitSum(sb.String(), k)
		}

		return _digitSum(s, k)
	}

	t.Log(digitSum("11111222223", 3))
	t.Log(digitSum("00000000", 3))
}
