package problem513

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Test513(t *testing.T) {
	treeNode1 := TreeNode{Val: 20}
	treeNode2 := TreeNode{Val: 6}
	//treeNode3 := TreeNode{Val: 4}
	//treeNode4 := TreeNode{Val: 3}
	treeNode5 := TreeNode{Val: 15, Left: &treeNode2, Right: &treeNode1}
	treeNode6 := TreeNode{Val: 5, Left: nil, Right: nil}
	treeNode7 := TreeNode{Val: 10, Left: &treeNode6, Right: &treeNode5}
	fmt.Println(findBottomLeftValue(&treeNode7))
}

func findBottomLeftValue(root *TreeNode) int {
	maxValue, _ := findBottomLeft(root, 0)
	return maxValue
}

func findBottomLeft(root *TreeNode, currentDepth int) (maxValue int, maxDepth int) {
	if root == nil {
		return 0, currentDepth
	}

	currentDepth++

	if root.Left == nil && root.Right == nil {
		return root.Val, currentDepth
	}

	var leftMaxValue int
	var leftMaxDepth int

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		leftMaxValue = root.Left.Val
		leftMaxDepth = currentDepth + 1
	} else {
		leftMaxValue, leftMaxDepth = findBottomLeft(root.Left, currentDepth)
	}

	rightMaxValue, rightMaxDepth := findBottomLeft(root.Right, currentDepth)

	if leftMaxDepth >= rightMaxDepth {
		return leftMaxValue, leftMaxDepth
	} else {
		return rightMaxValue, rightMaxDepth
	}
}
