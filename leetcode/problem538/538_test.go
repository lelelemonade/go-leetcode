package problem538

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Test538(t *testing.T) {
	fmt.Println(convertBST(nil))
}

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	convToGT(root, &sum)
	return root
}

func convToGT(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	convToGT(root.Right, sum)
	temp := root.Val
	*sum += temp
	root.Val = *sum
	convToGT(root.Left, sum)
}
