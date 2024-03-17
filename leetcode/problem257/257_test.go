package problem257

import (
	"fmt"
	"strconv"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Test257(t *testing.T) {
	treeNode1 := TreeNode{Val: 3}
	treeNode2 := TreeNode{Val: 4}
	treeNode3 := TreeNode{Val: 4}
	treeNode4 := TreeNode{Val: 3}
	treeNode5 := TreeNode{Val: 2, Left: &treeNode2, Right: &treeNode1}
	treeNode6 := TreeNode{Val: 2, Left: &treeNode4, Right: &treeNode3}
	treeNode7 := TreeNode{Val: 1, Left: &treeNode6, Right: &treeNode5}
	fmt.Println(binaryTreePaths(&treeNode7))
}

func binaryTreePaths(root *TreeNode) []string {
	var result []string
	binaryFindPath(root, []int{}, &result)
	return result
}

func binaryFindPath(root *TreeNode, alreadyCountNodes []int, results *[]string) {
	if root == nil {
		return
	}
	alreadyCountNodes = append(alreadyCountNodes, root.Val)

	if root.Left == nil && root.Right == nil {
		var resultString string
		for _, v := range alreadyCountNodes {
			resultString = resultString + strconv.Itoa(v) + "->"
		}
		*results = append(*results, resultString[:len(resultString)-2])
	}

	binaryFindPath(root.Left, alreadyCountNodes, results)
	binaryFindPath(root.Right, alreadyCountNodes, results)
}
