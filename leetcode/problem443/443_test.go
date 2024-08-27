package problem443

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func compress(chars []byte) int {
	resultChar:=make([]byte,0)
	dupCount:=1
	addToCompression := func(char byte){
		resultChar=append(resultChar, char)
		if dupCount!=1{
			resultChar=append(resultChar, []byte(strconv.Itoa(dupCount))...)
		}
	}
	for i:=1;i<len(chars);i++{
		if chars[i]==chars[i-1]{
			dupCount++
		}else{
			addToCompression(chars[i-1])
			dupCount=1
		}
	}
	addToCompression(chars[len(chars)-1])

	copy(chars,resultChar)
	chars=chars[:len(resultChar)]

	return len(chars)
}

func Test443(t *testing.T) {
	assert.Equal(t,6,compress([]byte{'a','a','b','b','c','c','c'}))
}