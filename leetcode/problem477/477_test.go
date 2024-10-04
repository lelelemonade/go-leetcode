package problem477

func totalHammingDistance(nums []int) int {
    n:=len(nums)
	result:=0
	for i:=0;i<30;i++{
		totalDiff:=0
		for _,num:=range nums{
			totalDiff+=num>>i&1
		}
		result+=totalDiff*(n-totalDiff)
	}
	return result
}