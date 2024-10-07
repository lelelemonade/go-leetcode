package problem2696

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func minLength(s string) int {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])

		if len(stack) < 2 {
			continue
		}

		if (stack[len(stack)-2] == 'A' && stack[len(stack)-1] == 'B') || (stack[len(stack)-2] == 'C' && stack[len(stack)-1] == 'D') {
			stack = stack[:len(stack)-2]
		}
	}
	return len(stack)
}

func Test2696(t *testing.T){
	assert.Equal(t,2,minLength("ABFCACDB"))
}
