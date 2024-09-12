package problem1343

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func numOfSubarrays(arr []int, k int, threshold int) int {
    total:=0
	for i:=0;i<k;i++{
		total+=arr[i]
	}

	result:=0

	for i:=k;i<len(arr);i++{
		if total>=threshold*k{
			result++
		}
		total+=arr[i]
		total-=arr[i-k]
	}

	if total>=threshold*k{
		result++
	}

	return result
}

func Test1343(t *testing.T) {
	assert.Equal(t,3,numOfSubarrays([]int{2,2,2,2,5,5,5,8},3,4))
}