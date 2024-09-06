package problem543

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	result:=0
	var dfs func(start *TreeNode) int
	dfs = func(start *TreeNode)int{
		if start == nil {
			return 0
		}
		leftLength:=dfs(start.Left)
		rightLength:=dfs(start.Right)
		result=max(result,leftLength+rightLength)
		return max(leftLength,rightLength)+1
	}
	dfs(root)
	return result
}
