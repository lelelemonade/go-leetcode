package problem2576

import "slices"

func maxNumOfMarkedIndices(nums []int) int {
    slices.Sort(nums)
	result:=0
	left,right:=0,len(nums)/2
	for right < len(nums)&&left<len(nums)/2 {
		for 2*nums[left]>nums[right]{
			right++
			if right>=len(nums){
				return result
			}
		}
		result+=2
		left++
		right++
	}
	return result
}