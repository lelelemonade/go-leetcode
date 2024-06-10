package problem27

import (
	"fmt"
	"testing"
)

func Test_Problem27(t *testing.T) {
	testArray := []int{0, 4, 4, 0, 4, 4, 4, 0, 2}

	fmt.Println(removeElement(testArray, 4))
	fmt.Println(testArray)
}

func removeElement(nums []int, val int) int {
	left, right := 0, 0

	for right != len(nums) {
		if nums[right] == val {
			right++
		} else {
			nums[left] = nums[right]
			left++
			right++
		}
	}

	return left
}
