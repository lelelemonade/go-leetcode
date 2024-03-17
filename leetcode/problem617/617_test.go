package problem617

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Test617(t *testing.T) {
	treeNode1 := TreeNode{Val: 5}
	treeNode2 := TreeNode{Val: 2}
	treeNode3 := TreeNode{Val: 3, Left: &treeNode1, Right: nil}
	treeNode4 := TreeNode{Val: 1, Left: &treeNode3, Right: &treeNode2}

	treeNode5 := TreeNode{Val: 7}
	treeNode6 := TreeNode{Val: 4}
	treeNode7 := TreeNode{Val: 3, Left: nil, Right: &treeNode5}
	treeNode8 := TreeNode{Val: 1, Left: nil, Right: &treeNode6}
	treeNode9 := TreeNode{Val: 2, Left: &treeNode8, Right: &treeNode7}
	fmt.Println(mergeTrees(&treeNode4, &treeNode9))
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 == nil && root2 != nil {
		return root2
	}

	if root1 != nil && root2 == nil {
		return root1
	}

	if root1 != nil && root2 != nil {
		root1.Val = root1.Val + root2.Val
	}

	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)

	return root1
}
