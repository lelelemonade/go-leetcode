package problem118

import (
	"fmt"
	"testing"
)

func generate(numRows int) [][]int {
    results:=make([][]int,0)
	results = append(results, []int{1})

	for i:=1;i<numRows;i++{
		result:=[]int{1}
		for j:=1;j<i;j++{
			result = append(result, results[i-1][j-1]+results[i-1][j])
		}
		result = append(result, 1)
		results = append(results, result)
	}

	return results
}

func Test118(t *testing.T) {
	fmt.Println(generate(5))
}