package problem437

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	result:=0
	var dfs func(start *TreeNode,pathSum []int)
	dfs = func(start *TreeNode,pathSum []int)  {
		if start==nil{
			return
		}

		pathSum=append(pathSum, 0)
		for i:=0;i<len(pathSum);i++{
			pathSum[i]+=start.Val
			if pathSum[i]==targetSum{
				result++
			}
		}

		currentPathSum:=make([]int,len(pathSum))

		copy(currentPathSum,pathSum)

		dfs(start.Left,pathSum)
		dfs(start.Right,currentPathSum)
	}

	dfs(root,[]int{})

	return result
}

func Test437(t *testing.T) {
	node1:=&TreeNode{Val:1,Left:nil,Right: nil}
	node2:=&TreeNode{Val:5,Left:nil,Right: nil}
	node3:=&TreeNode{Val:2,Left:nil,Right: nil}
	node4:=&TreeNode{Val:7,Left:nil,Right: nil}
	node5:=&TreeNode{Val:4,Left:nil,Right: nil}
	node6:=&TreeNode{Val:13,Left:node2,Right: node1}
	node7:=&TreeNode{Val:11,Left:node4,Right: node3}
	node8:=&TreeNode{Val:8,Left:node6,Right: node5}
	node9:=&TreeNode{Val:4,Left:node7,Right: nil}
	node10:=&TreeNode{Val:5,Left:node9,Right: node8}

	assert.Equal(t,3,pathSum(node10,22))
}