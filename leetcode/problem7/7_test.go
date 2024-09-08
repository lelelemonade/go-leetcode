package problem7

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func reverse(x int) int {
	upperBound:=math.MaxInt32

	internalReverse := func (a int) int{
		nextX:=0
		for a!=0{
			if nextX*10>=upperBound{
				return 0
			}
			nextX*=10
			if nextX+a%10>=upperBound{
				return 0
			}
			nextX+=a%10
			a/=10
		}
		return nextX
	}
    if x<0{
		return -internalReverse(-x)
	}

	return internalReverse(x)
}

func Test7(t *testing.T) {
	assert.Equal(t,321,reverse(123))
}