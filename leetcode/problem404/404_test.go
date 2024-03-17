package problem404

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var leftResult int

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		leftResult = root.Left.Val
	} else {
		leftResult = sumOfLeftLeaves(root.Left)
	}

	rightResult := sumOfLeftLeaves(root.Right)
	return leftResult + rightResult
}
