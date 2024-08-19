package problem114

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	flattenDFS(root)
}

func flattenDFS(root *TreeNode)(tail *TreeNode) {
	if root == nil || (root.Left==nil && root.Right== nil) {
		return root
	}
	if root.Right == nil {
		root.Right=root.Left
		root.Left=nil
		return flattenDFS(root.Right)
	}
	if root.Left==nil {
		return flattenDFS(root.Right)
	}
	right := root.Right
	root.Right=root.Left
	root.Left=nil
	flattenDFS(root.Right).Right=right
	return flattenDFS(right)
}
