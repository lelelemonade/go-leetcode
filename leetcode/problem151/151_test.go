package problem151

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)


func reverseWords(s string) string {
    sRune := []rune(strings.TrimSpace(s))

	reverseRuneArray(sRune)

	slow,fast:=0,0
	for ;slow<len(sRune)&&fast<len(sRune);{
		if sRune[fast]!=' ' || sRune[slow-1]!=' '{
			sRune[slow]=sRune[fast]
			slow++
		}
		fast++
	}
	sRune=sRune[:slow]

	for slow,fast:=0,0;slow<len(sRune)&&fast<=len(sRune);{
		if fast==len(sRune) || sRune[fast] == ' ' {
			reverseRuneArray(sRune[slow:fast])
			slow=fast+1
		}
		fast++
	}

	return string(sRune)
}

func reverseRuneArray(sRune []rune){
	for i := 0; i<len(sRune)/2;i++{
		sRune[i],sRune[len(sRune)-1-i]=sRune[len(sRune)-1-i],sRune[i]
	}
}

func Test151(t *testing.T) {
	// assert.Equal(t,"world hello",reverseWords("  hello world  "))
	assert.Equal(t,"example good a",reverseWords("a good   example"))
}