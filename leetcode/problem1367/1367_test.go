package problem1367

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}
	var dfs func(h *ListNode, r *TreeNode)bool
	dfs=func(h *ListNode, r *TreeNode)bool{
		if h==nil{
			return true
		}
		if r==nil{
			return false
		}
		if h.Val!=r.Val{
			return false
		}
		
		return dfs(h.Next,r.Left)||dfs(h.Next,r.Right)
	}

	return dfs(head,root)||isSubPath(head,root.Left)||isSubPath(head,root.Right)
}

func Test1367(t *testing.T) {
	head:=&ListNode{Val: 1,Next: &ListNode{Val:10}}
	treeNode1:=&TreeNode{Val: 1,Left: nil,Right:nil}
	treeNode2:=&TreeNode{Val: 9,Left: nil,Right:nil}
	treeNode4:=&TreeNode{Val: 10,Left: treeNode2,Right:nil}
	treeNode3:=&TreeNode{Val: 1,Left: treeNode4,Right:treeNode1}
	root:=&TreeNode{Val: 1,Left: nil,Right:treeNode3}
	assert.Equal(t,true,isSubPath(head,root))
}