package problem2958

func maxSubarrayLength(nums []int, k int) int {
    slidingWindowNumCount:=make(map[int]int)
	left:=0
	right:=0
	result:=0
	for right!=len(nums){
		slidingWindowNumCount[nums[right]]++
		for slidingWindowNumCount[nums[right]]>k{
			slidingWindowNumCount[nums[left]]--
			left++
		}
		result=max(result,right-left+1)
		right++
	}
	return result
}