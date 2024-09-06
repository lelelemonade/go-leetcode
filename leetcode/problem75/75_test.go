package problem75

func sortColors(nums []int)  {
	zeroPointer,onePointer:=0,0

	for i:=0;i<len(nums);i++{
		if nums[i]==0{
			nums[zeroPointer],nums[i]=nums[i],nums[zeroPointer]
			if zeroPointer<onePointer{
				nums[i], nums[onePointer] = nums[onePointer], nums[i]
			}
			zeroPointer++
			onePointer++
		}else if nums[i]==1{
			nums[onePointer],nums[i]=nums[i],nums[onePointer]
			onePointer++
		}
	}
}