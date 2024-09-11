package problem2220

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func minBitFlips(start int, goal int) int {
    bitSlice:=make([]bool,0)

	for start!=0{
		bitSlice=append(bitSlice, start%2==1)
		start>>=1
	}

	result:=0
	i:=0

	for goal!=0||i<len(bitSlice){
		if i>=len(bitSlice){
			if goal%2==1{
				result++
			}
		}else if goal%2==1!=bitSlice[i]{
			result++
		}
		i++
		goal>>=1
	}

	return result
}

func Test2220(t *testing.T) {
	// assert.Equal(t,3,minBitFlips(10,7))
	assert.Equal(t,3,minBitFlips(10,82))
}