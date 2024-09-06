package problem1800

func maxAscendingSum(nums []int) int {
	result:=nums[0]
	sum:=nums[0]
    for i:=1;i<len(nums);i++{
		if nums[i]>nums[i-1]{
			sum+=nums[i]
		}else{
			sum=nums[i]
		}
		result=max(result,sum)
	}
	return result
}