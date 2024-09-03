package problem69

func mySqrt(x int) int {
	left,right:=0,x

    for left<=right{
		mid:=(right-left)/2
		pow:=mid*mid
		nextPow:=(mid+1)*(mid+1)
		if pow<=x&&nextPow>x{
			return mid
		}else if pow<x{
			right=mid-1
		}else{
			left=mid+1
		}
	}
	return 1
}