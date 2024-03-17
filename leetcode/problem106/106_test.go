package problem106

import (
	"fmt"
	"slices"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Test106(t *testing.T) {
	root := buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	fmt.Println(root)
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	postRootIndex := slices.Index(inorder, postorder[len(postorder)-1])

	leftNode := buildTree(inorder[:postRootIndex], postorder[:postRootIndex])

	rightNode := buildTree(inorder[postRootIndex+1:], postorder[postRootIndex:len(postorder)-1])

	root := TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  leftNode,
		Right: rightNode,
	}

	return &root
}
