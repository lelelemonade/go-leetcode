package problem268

func missingNumber(nums []int) int {
    sum:=0
	for _,v:=range nums{
		sum+=v
	}

	return (len(nums)*(len(nums)+1))/2-sum
}