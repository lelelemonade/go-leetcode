package problem74

func searchMatrix(matrix [][]int, target int) bool {
    m,n:=len(matrix),len(matrix[0])
	left:=0
	right:=m*n-1
	mid:=(left+right)/2
	for left<=right{
		midX:=mid/n
		midY:=mid%n
		if matrix[midX][midY]==target{
			return true
		}else if matrix[midX][midY]<target{
			left=mid+1
		}else if matrix[midX][midY]>target{
			right=mid-1
		}
		mid=(left+right)/2
	}
	return false
}