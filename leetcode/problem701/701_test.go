package problem701

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	dfs(root, val)

	return root
}

func dfs(root *TreeNode, val int) {
	if val < root.Val {
		if root.Left == nil {
			root.Left = &TreeNode{
				Val: val,
			}
		} else {
			dfs(root.Left, val)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{
				Val: val,
			}
		} else {
			dfs(root.Right, val)
		}
	}
}
