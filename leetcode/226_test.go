package leetcode

import (
	"fmt"
	"testing"
)

func Test226(t *testing.T) {
	treeNode1 := TreeNode{Val: 1}
	treeNode2 := TreeNode{Val: 3}
	treeNode3 := TreeNode{Val: 6}
	treeNode4 := TreeNode{Val: 9}
	treeNode5 := TreeNode{Val: 2, Left: &treeNode2, Right: &treeNode1}
	treeNode6 := TreeNode{Val: 7, Left: &treeNode4, Right: &treeNode3}
	treeNode7 := TreeNode{Val: 4, Left: &treeNode6, Right: &treeNode5}
	result := invertTree(&treeNode7)
	fmt.Println(*result)
}

func invertTree(root *TreeNode) *TreeNode {
	if nil == root {
		return nil
	}

	root.Right, root.Left = root.Left, root.Right
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
