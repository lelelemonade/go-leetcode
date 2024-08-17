package problem224

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func calculate(s string) int {
	stack := make([]CalValue, 0)

	shrunkStack := func() bool {
		stackLen := len(stack)

		if stackLen==2 && stack[0].calc=='-'&&stack[1].value!=0{
			stack[stackLen-2].calc=0
			stack[stackLen-2].value=- stack[stackLen-1].value
			stack = stack[:stackLen-1]
			return true
		}

		if stackLen < 3 {
			return false
		}

		switch stack[stackLen-2].calc {
		case '+':
			stack[stackLen-3].value = stack[stackLen-3].value + stack[stackLen-1].value
			stack = stack[:stackLen-2]
			return true
		case '-':
			if stack[stackLen-3].calc=='('{
				stack[stackLen-2].calc=0
				stack[stackLen-2].value=- stack[stackLen-1].value
				stack = stack[:stackLen-1]
				return true
			}
			stack[stackLen-3].value = stack[stackLen-3].value - stack[stackLen-1].value
			stack = stack[:stackLen-2]
			return true
		case '/':
			stack[stackLen-3].value = stack[stackLen-3].value / stack[stackLen-1].value
			stack = stack[:stackLen-2]
			return true
		case '*':
			stack[stackLen-3].value = stack[stackLen-3].value * stack[stackLen-1].value
			stack = stack[:stackLen-2]
			return true
		default:
			return false
		}
	}

	for i, v := range s {
		switch v {
		case '(', '+', '-', '/', '*':
			stack = append(stack, CalValue{calc: v})
		case ' ':
			continue
		case ')':
			stackLen := len(stack)
			stack[stackLen-2] = stack[stackLen-1]
			stack = stack[:stackLen-1]
			for (i==len(s)-1||(s[i+1]>='0'&&s[i+1]<='9'))&& shrunkStack() {

			}
		default:
			if len(stack) > 0 && stack[len(stack)-1].value != 0 {
				stack[len(stack)-1].value = stack[len(stack)-1].value*10 + v - '0'
			}else{
				stack = append(stack, CalValue{value: v - '0'})
			}
			
			for (i==len(s)-1||(s[i+1]<'0'&&s[i+1]>'9'))&& shrunkStack() {

			}
		}
	}
	return int(stack[0].value)
}

type CalValue struct {
	calc  rune
	value int32
}

func Test224(t *testing.T) {
	assert.Equal(t, 23, calculate("(1+(4+5+2)-3)+(6+8)"))
	// assert.Equal(t, 2, calculate("1 + 1"))
	// assert.Equal(t, 2147483647, calculate("2147483647"))
	// assert.Equal(t, 3, calculate("1-(     -2)"))
	// assert.Equal(t, -10, calculate("1-11"))
}
