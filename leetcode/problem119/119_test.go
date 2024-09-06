package problem119

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getRow(rowIndex int) []int {
	result:=[]int{1}

	for i:=1;i<=rowIndex;i++{
		lastResultJ:=1
		for j:=1;j<i;j++{
			nextValue:=result[j]+lastResultJ
			lastResultJ=result[j]
			result[j]=nextValue
		}

		result = append(result, 1)
	}

	return result
}

func Test119(t *testing.T) {
	assert.Equal(t,[]int{1},getRow(0))
	assert.Equal(t,[]int{1,1},getRow(1))
	assert.Equal(t,[]int{1,2,1},getRow(2))
	assert.Equal(t,[]int{1,3,3,1},getRow(3))
}