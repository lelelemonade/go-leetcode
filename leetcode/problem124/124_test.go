package problem124

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	result:=math.MinInt
	maxPathSumDFS(root,&result)
	return result
}

func maxPathSumDFS(root *TreeNode, result *int)int{
	if root == nil {
		return 0
	}
	if root.Left==nil && root.Right== nil {
		*result=max(*result,root.Val)
		return root.Val
	}

	left:=maxPathSumDFS(root.Left,result)
	right:=maxPathSumDFS(root.Right,result)

	*result=max(*result,left+right+root.Val,left+root.Val,right+root.Val,root.Val)

	return max(left,right,0)+root.Val
}

// func maxPathSumDFS(root *TreeNode, result *int)(chooseLeft int, chooseRight int){
// 	if root == nil {
// 		return 0,0
// 	}
// 	if root.Left==nil && root.Right== nil {
// 		return root.Val,root.Val
// 	}

// 	leftChooseLeft,leftChooseRight:=maxPathSumDFS(root.Left,result)
// 	rightChooseLeft,rightChooseRight:=maxPathSumDFS(root.Right,result)

// 	current:=max(maxPathSum(root.Left),maxPathSum(root.Right))+root.Val

// 	*result=max(*result,current)

// 	return current
// }