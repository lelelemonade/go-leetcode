package problem2574

func leftRightDifference(nums []int) []int {
    total:=0
	for _,v:=range nums{
		total+=v
	}

	leftSum:=0
	result:=make([]int,len(nums))

	for i:=0;i<len(nums);i++{
		rightSum:=total-leftSum-nums[i]

		result[i]=abs(leftSum-rightSum)
		leftSum+=nums[i]
	}
	return result
}

func abs(a int) int {
	if a>0{
		return a
	}
	return -a
}