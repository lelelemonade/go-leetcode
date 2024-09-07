package problem589

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	result:=make([]int,0)
	var dfs func(start *Node)
	dfs=func(start *Node){
		if start == nil {
			return
		}
		result = append(result, start.Val)
		for _,v:=range start.Children{
			dfs(v)
		}
	}
	dfs(root)
	return result
}
