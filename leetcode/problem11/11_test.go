package problem11

func maxArea(height []int) int {
    left,right:=0,len(height)-1
	area:=min(height[left],height[right])*(len(height)-1)

	for left<right{
		heightLeft:=height[left]
		heightRight:=height[right]
		if heightLeft<heightRight{
			for height[left]<=heightLeft{
				left++
				if left>=right{
					return area
				}
			}
		}else{
			for height[right]<=heightRight{
				right--
				if left>=right{
					return area
				}
			}
		}
		area=max(area,min(height[left],height[right])*(right-left))
	}
	return area
}