package lcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test388(t *testing.T) {
	lengthLongestPath := func(input string) int {
		max := func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}

		paths := strings.Split(input, "\n")
		levelMax := make([]int, len(paths)+1)
		ans := 0
		for _, path := range paths {
			level := strings.Count(path, "\t")
			ln := len(path) - level

			preLevelLen := 0
			if level-1 >= 0 {
				preLevelLen = levelMax[level-1]
			}
			if strings.Contains(path, ".") {
				ans = max(ans, ln+preLevelLen)
			} else {
				levelMax[level] = preLevelLen + ln + 1
			}
		}
		return ans
	}

	assert.Equal(t, lengthLongestPath("dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"), 20)
	assert.Equal(t, lengthLongestPath("dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"), 32)
	assert.Equal(t, lengthLongestPath("a"), 0)
	assert.Equal(t, lengthLongestPath("file1.txt\nfile2.txt\nlongfile.txt"), 12)

}
