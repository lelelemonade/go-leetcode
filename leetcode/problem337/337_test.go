package problem337

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	dpMap := make(map[*TreeNode]int)

	dfs(root, dpMap)

	return dpMap[root]
}

func dfs(root *TreeNode, dpMap map[*TreeNode]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		dpMap[root] = root.Val
		return
	}

	dfs(root.Left, dpMap)
	dfs(root.Right, dpMap)

	left := 0
	right := 0
	leftLeft := 0
	leftRight := 0
	rightLeft := 0
	rightRight := 0

	if root.Left != nil {
		left = dpMap[root.Left]
		if root.Left.Left != nil {
			leftLeft = dpMap[root.Left.Left]
		}
		if root.Left.Right != nil {
			leftRight = dpMap[root.Left.Right]
		}
	}

	if root.Right != nil {
		right = dpMap[root.Right]
		if root.Right.Left != nil {
			rightLeft = dpMap[root.Right.Left]
		}
		if root.Right.Right != nil {
			rightRight = dpMap[root.Right.Right]
		}
	}

	dpMap[root] = max(
		left+right,
		leftLeft+
			leftRight+
			rightLeft+
			rightRight+
			root.Val,
	)
}
