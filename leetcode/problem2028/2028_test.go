package problem2028

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func missingRolls(rolls []int, mean int, n int) []int {
	sum:=0
	for _,v:=range rolls{
		sum+=v
	}
    missingCount:=mean*(n+len(rolls))-sum
	if missingCount<n||missingCount>n*6{
		return []int{}
	}
	avg:=missingCount/n
	remainder:=missingCount%n
	result:=make([]int,0)
	for i:=0;i<n;i++{
		if i<remainder{
			result = append(result, avg+1)
		}else{
			result = append(result, avg)
		}
	}
	return result
}

func Test2028(t *testing.T) {
	assert.Equal(t,[]int{6,6},missingRolls([]int{3,2,4,3},4,2))
}