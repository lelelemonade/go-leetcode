package problem739

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test739(t *testing.T) {
	testCase1 := []int{73, 74, 75, 71, 69, 72, 76, 73}
	assert.Equal(t, []int{1, 1, 4, 2, 1, 1, 0, 0}, dailyTemperatures(testCase1))

	testCase2 := []int{30, 40, 50, 60}
	assert.Equal(t, []int{1, 1, 1, 0}, dailyTemperatures(testCase2))
	testCase3 := []int{30, 60, 90}
	assert.Equal(t, []int{1, 1, 0}, dailyTemperatures(testCase3))

}

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	result := make([]int, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		if len(stack) == 0 || temperatures[stack[len(stack)-1]] >= temperatures[i] {
			stack = append(stack, i)
			continue
		}
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			result[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	return result
}
