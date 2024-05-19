package problem2963

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test2963(t *testing.T) {
	testCase1 := []int{1, 2, 3, 4}
	assert.Equal(t, 8, numberOfGoodPartitions(testCase1))

	testCase2 := []int{1, 1, 1, 1}
	assert.Equal(t, 1, numberOfGoodPartitions(testCase2))

	testCase3 := []int{1, 2, 1, 3}
	assert.Equal(t, 2, numberOfGoodPartitions(testCase3))

	testCase4 := []int{1, 1, 1, 3, 4}
	assert.Equal(t, 4, numberOfGoodPartitions(testCase4))

	testCase5 := []int{1, 5, 1, 5, 6}
	assert.Equal(t, 2, numberOfGoodPartitions(testCase5))

	testCase6 := []int{2, 3, 2, 8, 8}
	assert.Equal(t, 2, numberOfGoodPartitions(testCase6))

	testCase7 := make([]int, 100000)
	for i := 0; i < len(testCase7); i++ {
		testCase7[i] = i + 1
	}
	assert.Equal(t, 303861760, numberOfGoodPartitions(testCase7))

	testCase8 := make([]int, 100000)
	for i := 0; i < len(testCase8); i++ {
		testCase8[i] = 1
	}
	assert.Equal(t, 1, numberOfGoodPartitions(testCase8))
}

func numberOfGoodPartitions(nums []int) int {
	unionFind := make([]int, len(nums))
	numStartIds := make(map[int]int)

	for i := 0; i < len(unionFind); i++ {
		unionFind[i] = i
	}

	for i := 0; i < len(nums); i++ {
		if _, e := numStartIds[nums[i]]; !e {
			numStartIds[nums[i]] = i
		}
	}

	for i := 0; i < len(nums); i++ {
		for j := numStartIds[nums[i]]; j <= i; j++ {
			unionFind[j] = unionFind[numStartIds[nums[i]]]
		}

		numStartIds[nums[i]] = i
	}

	result := 1
	uniqueUnion := make(map[int]struct{})

	for i := 0; i < len(unionFind); i++ {
		uniqueUnion[unionFind[i]] = struct{}{}
	}

	for i := 0; i < len(uniqueUnion)-1; i++ {
		result = (result * 2) % 1000000007
	}

	return result
}
