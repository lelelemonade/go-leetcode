package problem230

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test230Case1(t *testing.T) {
	node1 := TreeNode{Val: 1, Left: nil, Right: nil}
	node2 := TreeNode{Val: 2, Left: &node1, Right: nil}
	node3 := TreeNode{Val: 4, Left: nil, Right: nil}
	node4 := TreeNode{Val: 3, Left: &node2, Right: &node3}
	node5 := TreeNode{Val: 6, Left: nil, Right: nil}
	node6 := TreeNode{Val: 5, Left: &node4, Right: &node5}

	assert.Equal(t, kthSmallest(&node6, 3), 3)
}

func Test230Case2(t *testing.T) {
	node1 := TreeNode{Val: 2, Left: nil, Right: nil}
	node2 := TreeNode{Val: 1, Left: nil, Right: &node1}
	node3 := TreeNode{Val: 4, Left: nil, Right: nil}
	node4 := TreeNode{Val: 3, Left: &node2, Right: &node3}

	assert.Equal(t, kthSmallest(&node4, 1), 1)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		stack, root = stack[:len(stack)-1], stack[len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}
