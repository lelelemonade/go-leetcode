package problem112

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return dfs(root, 0, targetSum)
}

func dfs(root *TreeNode, currentSum int, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && root.Val+currentSum == targetSum {
		return true
	}

	if dfs(root.Left, currentSum+root.Val, targetSum) {
		return true
	}

	if dfs(root.Right, currentSum+root.Val, targetSum) {
		return true
	}

	return false
}
