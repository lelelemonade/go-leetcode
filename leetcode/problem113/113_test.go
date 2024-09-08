package problem113

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	results:=make([][]int,0)
	var dfs func(nodes []int,total int, start *TreeNode)
	dfs=func(nodes []int,total int,start *TreeNode){
		if start==nil{
			return
		}
		if total+start.Val==targetSum&&start.Left==nil&&start.Right==nil{
			result:=make([]int,len(nodes)+1)
			copy(result,append(nodes,start.Val))
			results=append(results, result)
		}
		dfs(append(nodes,start.Val),total+start.Val,start.Left)
		dfs(append(nodes,start.Val),total+start.Val,start.Right)
	}
	dfs([]int{},0,root)
	return results
}
