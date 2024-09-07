package problem275

func hIndex(citations []int) int {
	left:=0
	right:=len(citations)-1
	for left<=right{
		mid:=(left+right)/2

		if citations[mid]>=len(citations)-mid{
			right=mid-1
		}else{
			left=mid+1
		}
	}
	return len(citations)-left
}