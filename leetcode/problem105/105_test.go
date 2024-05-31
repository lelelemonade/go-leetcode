package problem105

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootIndexOfInOrder := slices.Index(inorder, preorder[0])

	leftNode := buildTree(preorder[1:rootIndexOfInOrder+1], inorder[:rootIndexOfInOrder])
	rightNode := buildTree(preorder[rootIndexOfInOrder+1:], inorder[rootIndexOfInOrder+1:])

	return &TreeNode{
		Val:   preorder[0],
		Left:  leftNode,
		Right: rightNode,
	}
}
