package problem2966

import "slices"

func divideArray(nums []int, k int) [][]int {
    slices.Sort(nums)

	result:=make([][]int,0)

	for i:=0;i<=len(nums)-3;i+=3{
		if nums[i+2]-nums[i]>k{
			return [][]int{}
		}
		result=append(result, nums[i:i+3])
	}

	return result
}