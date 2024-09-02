package problem1318

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func minFlips(a int, b int, c int) int {
    result:=0
	for a!=0||b!=0||c!=0{
		aBit:=a%2
		bBit:=b%2
		cBit:=c%2
		if cBit==1{
			if aBit!=1&&bBit!=1{
				result++
			}
		}else{
			if aBit!=0{
				result++
			}
			if bBit!=0{
				result++
			}
		}
		a>>=1
		b>>=1
		c>>=1
	}
	return result
}

func Test1318(t *testing.T) {
	assert.Equal(t,3,minFlips(8,3,5))
	
}