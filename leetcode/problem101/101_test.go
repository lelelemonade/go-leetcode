package problem101

import (
	"fmt"
	"testing"
)

func Test101(t *testing.T) {
	treeNode1 := TreeNode{Val: 3}
	treeNode2 := TreeNode{Val: 4}
	treeNode3 := TreeNode{Val: 4}
	treeNode4 := TreeNode{Val: 3}
	treeNode5 := TreeNode{Val: 2, Left: &treeNode2, Right: &treeNode1}
	treeNode6 := TreeNode{Val: 2, Left: &treeNode4, Right: &treeNode3}
	treeNode7 := TreeNode{Val: 1, Left: &treeNode6, Right: &treeNode5}
	fmt.Println(isSymmetric(&treeNode7))
}

func isSymmetric(root *TreeNode) bool {
	if nil == root {
		return true
	}

	var queue []*TreeNode

	queue = append(queue, root.Left)
	queue = append(queue, root.Right)

	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		if left == nil && right == nil {
			queue = queue[2:]
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left, right.Right, left.Right, right.Left)
		queue = queue[2:]
	}

	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
