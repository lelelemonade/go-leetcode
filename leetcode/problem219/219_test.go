package problem219

func containsNearbyDuplicate(nums []int, k int) bool {
    slidingWindowIdx:=make(map[int]struct{})
	
	for i:=0;i<len(nums);i++ {
		if len(slidingWindowIdx)>k {
			delete(slidingWindowIdx,nums[i-k-1])
		}
		if _,e:=slidingWindowIdx[nums[i]];!e{
			slidingWindowIdx[nums[i]]=struct{}{}
		}else{
			return true
		}
	}

	return false
}