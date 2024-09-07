package problem594

import "slices"

func findLHS(nums []int) int {
    slices.Sort(nums)

	result:=0
	for i:=0;i<len(nums);i++{
		cnt:=0
		resultActivate:=false
		for j:=i;j<len(nums);j++{
			if nums[j]-nums[i]<1{
				cnt++
			}else if nums[j]-nums[i]==1 {
				resultActivate=true
				cnt++
			}else{
				break
			}
		}
		if resultActivate{
			result=max(result,cnt)
		}
	}
	return result
}