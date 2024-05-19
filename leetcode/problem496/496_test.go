package problem496

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test486(t *testing.T) {
	//testCase1Nums1 := []int{4, 1, 2}
	//testCase1Nums2 := []int{1, 3, 4, 2}
	//testCase1Expect := []int{-1, 3, -1}
	//
	//assert.Equal(t, testCase1Expect, nextGreaterElement(testCase1Nums1, testCase1Nums2))

	testCase2Nums1 := []int{1, 3, 5, 2, 4}
	testCase2Nums2 := []int{5, 4, 3, 2, 1}
	testCase2Expect := []int{-1, -1, -1, -1, -1}

	assert.Equal(t, testCase2Expect, nextGreaterElement(testCase2Nums1, testCase2Nums2))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := make([]int, 0)
	resultMap := make(map[int]int)

	for i := 0; i < len(nums2); i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}

		if nums2[stack[len(stack)-1]] >= nums2[i] {
			stack = append(stack, i)
			continue
		}

		for len(stack) > 0 && nums2[stack[len(stack)-1]] < nums2[i] {
			resultMap[nums2[stack[len(stack)-1]]] = nums2[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	result := make([]int, len(nums1))
	for i, v := range nums1 {
		if greaterElem, exist := resultMap[v]; exist {
			result[i] = greaterElem
		} else {
			result[i] = -1
		}
	}

	return result
}
