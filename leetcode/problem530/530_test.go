package problem530

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	inorder := make([]int, 0)
	dfs(root, &inorder)

	minimumDifference := math.MaxInt
	for i := 0; i < len(inorder)-1; i++ {
		minimumDifference = min(minimumDifference, inorder[i+1]-inorder[i])
	}

	return minimumDifference
}

func dfs(root *TreeNode, inorder *[]int) {
	if root == nil {
		return
	}

	dfs(root.Left, inorder)
	*inorder = append(*inorder, root.Val)
	dfs(root.Right, inorder)
}
