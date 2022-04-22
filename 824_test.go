package lcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test824(t *testing.T) {
	toGoatLatin := func(sentence string) string {
		words := strings.Split(sentence, " ")
		var sb strings.Builder

		yuan := map[byte]bool{
			'a': true,
			'e': true,
			'i': true,
			'o': true,
			'u': true,
			'A': true,
			'E': true,
			'I': true,
			'O': true,
			'U': true,
		}

		target := "ma"

		for idx, word := range words {
			if _, ok := yuan[word[0]]; ok {
				//元音
				sb.WriteString(word)
				sb.WriteString(target)
			} else {
				sb.WriteString(word[1:])
				sb.WriteByte(word[0])
				sb.WriteString(target)
			}
			sb.WriteString(strings.Repeat("a", idx+1))
			if idx < len(words)-1 {
				sb.WriteByte(' ')
			}
		}

		return sb.String()
	}

	assert.Equal(t, toGoatLatin("I speak Goat Latin"), "Imaa peaksmaaa oatGmaaaa atinLmaaaaa")
	assert.Equal(t, toGoatLatin("The quick brown fox jumped over the lazy dog"), "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa")

}
