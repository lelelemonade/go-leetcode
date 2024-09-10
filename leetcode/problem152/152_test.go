package problem152

func maxProduct(nums []int) int {
	if len(nums)==1{
		return nums[0]
	}
	result:=nums[0]
	lastMax:=nums[0]
	lastMin:=nums[0]
	for i:=1;i<len(nums);i++{
		thisMax:=max(lastMax*nums[i],lastMin*nums[i],nums[i])
		result=max(thisMax,result)
		lastMin=min(lastMax*nums[i],lastMin*nums[i],nums[i])
		lastMax=thisMax
	}
	return result
}