package problem394

import (
	"slices"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func decodeString(s string) string {
	reverseString := func(s string) string {
		runeArray := []rune(s)
		slices.Reverse(runeArray)
		return string(runeArray)
	}
	s = reverseString(s)
	stack := make([]byte, 0)

	for i := 0; i < len(s);  {
		if s[i] == ']' {
			stack = append(stack, s[i])
			i++
		} else if s[i] == '[' {
			i++
			numArray := make([]byte, 0)
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				numArray = append(numArray, s[i])
				i++
			}
			n, e := strconv.Atoi(reverseString(string(numArray)))
			if e != nil {
				panic(e)
			}
			j:=len(stack)-1
			for j>=0&&stack[j]!=']'{
				j--
			}
			stack=append(stack[:j], stack[j+1:]...)
			lenStack:=len(stack)
			for k := 0; k < n-1; k++ {
				stack = append(stack, stack[j:lenStack]...)
			}
		}else{
			stack = append(stack, s[i])
			i++
		}
	}

	return reverseString(string(stack))
}

func Test394(t *testing.T) {
	assert.Equal(t, "aaabcbc", decodeString("3[a]2[bc]"))
	assert.Equal(t, "accaccacc", decodeString("3[a2[c]]"))
	assert.Equal(t, "abcabccdcdcdef", decodeString("2[abc]3[cd]ef"))
}
