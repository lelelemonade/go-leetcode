package problem95

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	result:=make([]*TreeNode,0)

	var constructBST func(candidates []int)*TreeNode
	constructBST = func(candidates []int)*TreeNode{
		if len(candidates)==0{
			return nil
		}

		for i,v:=range candidates{
			constructBST
		}
	}

	candidates:=make([]int,n)

	for i:=0;i<len(candidates);i++{
		candidates[i]=i+1
	}
	
	constructBST(candidates)

	return result
}
