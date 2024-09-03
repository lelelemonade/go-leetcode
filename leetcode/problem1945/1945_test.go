package problem1945

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getLucky(s string, k int) int {
	convert := ""
	for _, v := range s {
		convert += strconv.Itoa(int(byte(v)-'a') + 1)
	}
	for i := 0; i < k; i++ {
		next := 0
		for _, v := range convert {
			next += int(byte(v) - '0')
		}
		convert = strconv.Itoa(next)
	}
	result, _ := strconv.Atoi(convert)
	return result
}

func Test1945(t *testing.T) {
	assert.Equal(t, 8, getLucky("fleyctuuajsr", 5))
}
