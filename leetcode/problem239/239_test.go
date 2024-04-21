package problem239

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test239(t *testing.T) {
	testCase1 := []int{1, 3, -1, -3, 5, 3, 6, 7}
	expected1 := []int{3, 3, 5, 5, 6, 7}

	assert.Equal(t, maxSlidingWindow(testCase1, 3), expected1)

	testCase2 := []int{1}
	expected2 := []int{1}

	assert.Equal(t, maxSlidingWindow(testCase2, 1), expected2)

	testCase3 := []int{1, -1}
	expected3 := []int{1, -1}

	assert.Equal(t, maxSlidingWindow(testCase3, 1), expected3)

	testCase4 := []int{7, 2, 4}
	expected4 := []int{7, 4}

	assert.Equal(t, maxSlidingWindow(testCase4, 2), expected4)

	testCase5 := []int{1, 3, 1, 2, 0, 5}
	expected5 := []int{3, 3, 2, 5}

	assert.Equal(t, maxSlidingWindow(testCase5, 3), expected5)

	testCase6 := []int{9, 10, 9, -7, -4, -8, 2, -6}
	expected6 := []int{10, 10, 9, 2}

	assert.Equal(t, expected6, maxSlidingWindow(testCase6, 5))

}

func maxSlidingWindow(nums []int, k int) []int {
	monoQueue := list.New()
	result := make([]int, len(nums)-k+1)
	monoQueueMax := 0

	monoQueuePush := func(idx int) {
		if monoQueue.Len() > 0 && idx-monoQueue.Front().Value.(int) >= k {
			monoQueue.Remove(monoQueue.Front())
		}

		for monoQueue.Len() > 0 && nums[idx] > nums[monoQueue.Back().Value.(int)] {
			monoQueue.Remove(monoQueue.Back())
		}
		monoQueue.PushBack(idx)
		monoQueueMax = nums[monoQueue.Front().Value.(int)]
	}

	for i := 0; i < k; i++ {
		monoQueuePush(i)
	}

	result[0] = monoQueueMax

	for i := k; i < len(nums); i++ {
		monoQueuePush(i)
		result[i-k+1] = monoQueueMax
	}

	return result
}
