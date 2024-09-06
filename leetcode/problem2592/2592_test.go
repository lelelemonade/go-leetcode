package problem2592

import "slices"

func maximizeGreatness(nums []int) int {
    slices.Sort(nums)
	small,big:=len(nums)-1,len(nums)-1
	result:=0

	for small>=0{
		if nums[small]<nums[big]{
			big--
			result++
		}
		small--
	}
	return result
}