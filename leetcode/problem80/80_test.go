package problem80

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func removeDuplicates(nums []int) int {
	if len(nums)<=2 {
		return len(nums)
	}

	fast, slow := 2,2

	for fast < len(nums) {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}

func Test80(t *testing.T) {
	testCase1 := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(testCase1)
	result := removeDuplicates(testCase1)
	assert.Equal(t, result, 5)
	fmt.Println(testCase1)
}
