package problem104

import (
	"fmt"
	"runtime/debug"
	"testing"
)

func Test104(t *testing.T) {
	treeNode1 := TreeNode{Val: 3}
	treeNode2 := TreeNode{Val: 4}
	treeNode3 := TreeNode{Val: 4}
	treeNode4 := TreeNode{Val: 3}
	treeNode5 := TreeNode{Val: 2, Left: &treeNode2, Right: &treeNode1}
	treeNode6 := TreeNode{Val: 2, Left: &treeNode4, Right: &treeNode3}
	treeNode7 := TreeNode{Val: 1, Left: &treeNode6, Right: &treeNode5}
	fmt.Println(maxDepth(&treeNode7))
}

func maxDepth(root *TreeNode) int {
	maxDep := 0

	dfs(root, 0, &maxDep)

	return maxDep
}

func init() {
	debug.SetGCPercent(20)
}

func dfs(root *TreeNode, currentDepth int, maxDep *int) {
	if root == nil {
		return
	}
	if currentDepth+1 > *maxDep {
		*maxDep = currentDepth + 1
	}
	dfs(root.Left, currentDepth+1, maxDep)
	dfs(root.Right, currentDepth+1, maxDep)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
