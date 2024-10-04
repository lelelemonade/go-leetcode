package problem2433

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findArray(pref []int) []int {
    result:=make([]int,len(pref))
	result[0]=pref[0]
	resultPrefixXor:=pref[0]
	for i:=1;i<len(pref);i++{
		result[i]=pref[i]^resultPrefixXor
		resultPrefixXor^=result[i]
	}

	return result
}

func Test2433(t *testing.T) {
	assert.Equal(t,[]int{5,7,2,3,2},findArray([]int{5,2,0,3,1}))
}