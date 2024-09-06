package problem278

func isBadVersion(version int) bool

func firstBadVersion(n int) int {
    left:=1
	right:=n

	for left<=right{
		mid:=(left+right)/2

		isMidBad:=isBadVersion(mid)

		if isMidBad&&(mid==1||!isBadVersion(mid-1)){
			return mid
		}else if !isMidBad{
			left=mid+1
		}else{
			right=mid-1
		}
	}
	return 1
}