package problem2961

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getGoodIndices(variables [][]int, target int) []int {
	result := make([]int, 0)

	isGood := func(a, b, c, m int) bool {
		abMod:=a%10
		for i:=1;i<b;i++{
			abMod=(abMod*a)%10
		}
		cMod:=abMod%m
		for i:=1;i<c;i++{
			cMod=(cMod*abMod)%m
		}
		return cMod==target
	}

	for i := 0; i < len(variables); i++ {
		a := variables[i][0]
		b := variables[i][1]
		c := variables[i][2]
		m := variables[i][3]
		if isGood(a, b, c, m) {
			result = append(result, i)
		}
	}

	return result
}

func Test2961(t *testing.T){
	assert.Equal(t,[]int{0,2},getGoodIndices([][]int{{2,3,3,10},{3,3,3,1},{6,1,1,4}},2))
}
