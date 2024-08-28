package problem1448

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	var dfs func(start *TreeNode,maxofPath int,result *int)
	dfs =func(start *TreeNode,maxofPath int,result *int){
		if start==nil{
			return
		}
		if start.Val>=maxofPath{
			*result++
		}
		dfs(start.Left,max(maxofPath,start.Val),result)
		dfs(start.Right,max(maxofPath,start.Val),result)
	}

	var result int

	dfs(root,-10001,&result)

	return result
}
