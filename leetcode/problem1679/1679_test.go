package problem1679

import "slices"

func maxOperations(nums []int, k int) int {
    slices.Sort(nums)

	left,right:=0,len(nums)-1
	result:=0
	for left<right{
		if nums[left]+nums[right]<k{
			left++
		}else if nums[left]+nums[right]>k{
			right--
		}else{
			left++
			right--
			result++
		}
	}
	return result
}