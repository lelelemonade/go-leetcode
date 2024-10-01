package problem2595

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func evenOddBit(n int) []int {
    odd:=0
	even:=0

	i:=0
	for n!=0{
		i++
		if n%2==1{
			if i%2==1{
				odd++
			}else{
				even++
			}
		}
		n/=2
	}

	return []int{odd,even}
}

func Test2595(t *testing.T) {
	assert.Equal(t,[]int{1,2},evenOddBit(50))
	assert.Equal(t,[]int{0,1},evenOddBit(2))
}