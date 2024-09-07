package problem653

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	var dfs func(start *TreeNode,target *TreeNode)bool
	dfs=func(start *TreeNode,target *TreeNode)bool{
		if start == nil{
			return false
		}
		if target!=start&&k-target.Val == start.Val{
			return true
		}else if k-target.Val<start.Val{
			return dfs(start.Left,target)
		}else{
			return dfs(start.Right,target)
		}
	}
	queue:=make([]*TreeNode,1)
	queue[0]=root
	for len(queue)!=0{
		lenQueue:=len(queue)
		for i:=0;i<lenQueue;i++{
			if dfs(root,queue[i]){
				return true
			}
			if queue[i].Left!=nil{
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right!=nil{
				queue = append(queue, queue[i].Right)
			}
		}
		queue=queue[lenQueue:]
	}
	return false
}
