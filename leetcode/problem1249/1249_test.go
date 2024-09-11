package problem1249

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func minRemoveToMakeValid(s string) string {
    stack:=make([]int,0)
	for i:=0;i<len(s);i++{
		if s[i]==')'&&len(stack)>0&&s[stack[len(stack)-1]]=='('{
			stack=stack[:len(stack)-1]
		}else if s[i]=='('||s[i]==')'{
			stack=append(stack, i)
		}
	}
	for i:=len(stack)-1;i>=0;i--{
		s=s[:stack[i]]+s[stack[i]+1:]
	}
	return s
}

func Test1249(t *testing.T) {
	assert.Equal(t,"lee(t(c)o)de",minRemoveToMakeValid("lee(t(c)o)de)"))
}