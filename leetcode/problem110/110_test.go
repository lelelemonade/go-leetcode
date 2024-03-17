package problem110

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, balanced := isHeightBalance(root)
	return balanced
}

func isHeightBalance(root *TreeNode) (height int, balanced bool) {
	if nil == root {
		return 0, true
	}
	leftHeight, leftBalanced := isHeightBalance(root.Left)
	if !leftBalanced {
		return 0, false
	}
	rightHeight, rightBalanced := isHeightBalance(root.Right)
	if !rightBalanced {
		return 0, false
	}

	if rightHeight-leftHeight >= -1 && rightHeight-leftHeight <= 1 {
		return max(rightHeight, leftHeight) + 1, true
	}
	return 0, false
}
