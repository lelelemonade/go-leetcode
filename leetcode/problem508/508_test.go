package problem508

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	sumMap:=make(map[int]int)

	var dfs func(start *TreeNode)int
	dfs=func(start *TreeNode)int{
		if start == nil{
			return 0
		}
		left:=dfs(start.Left)
		right:=dfs(start.Right)
		sum:=left+right+start.Val
		sumMap[sum]++
		return sum
	}

	dfs(root)

	maxSum:=make([]int,0)
	maxCount:=0
	for k,v:=range sumMap{
		if v>maxCount{
			maxSum=[]int{k}
			maxCount=v
		}else if v==maxCount{
			maxSum=append(maxSum, k)
		}
	}
	return maxSum
}
