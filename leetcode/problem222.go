package leetcode

import "runtime/debug"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func init() {
	debug.SetGCPercent(20)
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := root.Left
	leftCount := 0

	for left != nil {
		left = left.Left
		leftCount++
	}

	right := root.Right
	rightCount := 0

	for right != nil {
		right = right.Right
		rightCount++
	}

	if leftCount == rightCount {
		return 2<<rightCount - 1
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1

}
