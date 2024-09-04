package problem93

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result:=make([]int,0)

	var dfs func(start *TreeNode)
	dfs = func(start *TreeNode){
		if start == nil {
			return
		}
		dfs(start.Left)
		result = append(result, start.Val)
		dfs(start.Right)
	}
	dfs(root)

	return result
}
