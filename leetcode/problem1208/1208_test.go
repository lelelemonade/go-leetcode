package problem1208

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func equalSubstring(s string, t string, maxCost int) int {
	abs:=func(a,b int)int{
		if a>b{
			return a-b
		}
		return b-a
	}
    diffTotal:=0
	left:=0
	maxLength:=0
	for right:=0;right<len(s);right++{
		diffTotal+=abs(int(s[right]),int(t[right]))
		for diffTotal>maxCost{
			diffTotal-=abs(int(s[left]),int(t[left]))
			left++
		}
		maxLength=max(maxLength,right-left+1)
	}
	return maxLength
}

func Test1208(t *testing.T) {
	assert.Equal(t,3,equalSubstring("abcd","bcdf",3))
	assert.Equal(t,2,equalSubstring("krrgw","zjxss",19))
}