package problem6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
    zigSlice:=make([][]byte,numRows)
	for i:=0;i<len(zigSlice);i++{
		zigSlice[i]=make([]byte,0)
	}

	for i:=0;i<len(s);i++{
		mod:=i%(2*numRows-2)
		if mod<numRows{
			zigSlice[mod]=append(zigSlice[mod], s[i])
		}else{
			zigSlice[numRows-2-mod%numRows]=append(zigSlice[numRows-2-mod%numRows], s[i])
		}
	}
	result:=""
	for _,v:=range zigSlice{
		result+=string(v)
	}
	return result
}

func Test6(t *testing.T) {
	assert.Equal(t,"PAHNAPLSIIGYIR",convert("PAYPALISHIRING",3))
	assert.Equal(t,"PINALSIGYAHRPI",convert("PAYPALISHIRING",4))
}