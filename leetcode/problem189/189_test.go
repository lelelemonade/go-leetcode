package problem189

import (
	"fmt"
	"testing"
)

func reverse(nums []int){
	for i := 0;i<len(nums)/2;i++ {
		nums[i],nums[len(nums)-1-i]=nums[len(nums)-1-i],nums[i]
	}
}

func rotate(nums []int, k int) {
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func Test189(t *testing.T) {
	testCase1 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(testCase1, 2)
	fmt.Println(testCase1)
}
