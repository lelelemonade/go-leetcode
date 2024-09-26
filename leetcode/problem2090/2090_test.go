package problem2090

func getAverages(nums []int, k int) []int {
    result:=make([]int,len(nums))
	rangeTotal:=0

	if len(nums)<2*k{
		for i:=0;i<len(nums);i++{
			result[i]=-1
		}
		return result
	}

	for i:=0;i<2*k;i++{
		rangeTotal+=nums[i]
	}

	for i,_:=range nums{
		if i< k||i>=len(nums)-k{
			result[i]=-1
		}else{
			rangeTotal+=nums[i+k]
			result[i]=rangeTotal/(2*k+1)
			rangeTotal-=nums[i-k]
		}
	}

	return result
}