package lcode

import (
	"strconv"
	"testing"
)

func Test43(t *testing.T) {

	multiply := func(num1 string, num2 string) string {
		var add func(num1 string, num2 string) string
		var mul func(num1 string, s string) string

		add = func(num1 string, num2 string) string {
			// 保证 num2 比 num1 更长
			if len(num1) > len(num2) {
				return add(num2, num1)
			}

			if num1 == "0" {
				return num2
			}
			if num2 == "0" {
				return num1
			}

			sb2 := []byte(num2)

			ln := len(num2)

			var jin byte = '0'

			for i := 0; i < ln; i++ {
				var n1 byte = '0'
				if len(num1)-i-1 >= 0 {
					n1 = num1[len(num1)-i-1]
				}

				n2 := sb2[len(num2)-i-1]

				t := n1 + n2 + jin - '0' - '0' - '0'
				if t >= 10 {
					// 需要进位
					jin = '1'

					sb2[len(num2)-i-1] = t + '0' - 10

				} else {
					sb2[len(num2)-i-1] = t + '0'
					jin = '0'
				}
			}
			if jin != '0' {
				sb2 = append([]byte{jin}, sb2...)
			}

			// 4
			//123456

			return string(sb2)
		}

		mul = func(num1 string, num2 string) string {
			if len(num1) < len(num2) {
				return mul(num2, num1)
			}

			if num2 == "10" {
				return num1 + "0"
			}

			if num1 == "10" {
				return num2 + "0"
			}

			if num1 == "0" || num2 == "0" {
				return "0"
			}

			if len(num1) == len(num2) && len(num1) == 1 {
				n1, _ := strconv.Atoi(num1)
				n2, _ := strconv.Atoi(num2)
				return strconv.Itoa(n1 * n2)
			}

			ans := "0"
			ln := len(num1)
			for i := 0; i < ln; i++ {
				n := string(num1[ln-i-1])
				t := mul(num2, n)
				for c := 0; c < i; c++ {
					t = mul(t, "10")
				}
				ans = add(ans, t)
			}

			return ans
		}

		ans := mul(num1, num2)
		flag := true
		for _, v := range ans {
			if v != '0' {
				flag = false
			}
		}

		if flag {
			return "0"
		}
		return ans

	}
	multiply("1", "2")

	// 21
	// 32
	// 21*2+30*21

	//t.Log(multiply("1", "2"))
	//t.Log(multiply("1", "21"))
	//t.Log(multiply("12", "34"))
	//t.Log(multiply("123", "456"))
	//t.Log(multiply("9133", "0"))
	//t.Log(multiply("103", "98"))
	t.Log(multiply("103", "98"))
	//t.Log(multiply("123456789", "987654321")) // 1,111,111,110

	//t.Log(multiply("401716832807512840963", "167141802233061013023557397451289113296441069"))
}
