package problem1372

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	result:=0
	var dfs func(start *TreeNode,lastDirection bool,zigzagCount int)
	dfs = func(start *TreeNode,lastDirection bool,zigzagCount int){
		if start==nil{
			return
		}
		result=max(result,zigzagCount)
		//last direction is left
		if lastDirection {
			dfs(start.Left,true,1)
			dfs(start.Right,false,zigzagCount+1)
		}else{
			dfs(start.Left,true,zigzagCount+1)
			dfs(start.Right,false,1)
		}
	}

	dfs(root,true,0)

	return result
}
