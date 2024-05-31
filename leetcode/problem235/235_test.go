package problem235

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val > q.Val && root.Val > p.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if root.Val < q.Val && root.Val < p.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
