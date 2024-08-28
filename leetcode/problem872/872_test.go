package problem872

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	root1Leaf:=""
	root2Leaf:=""

	var dfs func(root *TreeNode,result *string)

	dfs=func(root *TreeNode,result *string){
		if root==nil{
			return
		}
		if root.Left==nil && root.Right==nil{
			*result+=strconv.Itoa(root.Val)
			*result+="|"
		}
		dfs(root.Left,result)
		dfs(root.Right,result)
	}

	dfs(root1,&root1Leaf)
	dfs(root2,&root2Leaf)

	return root1Leaf==root2Leaf
}
