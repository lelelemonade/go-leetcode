package problem103

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := make([]*TreeNode,0)
	queue=append(queue, root)
	reverse:=false

	results:=make([][]int,0)

	for len(queue)!=0{
		lenQueue:=len(queue)
		result:=make([]int,0)
		if reverse{
			for i:=lenQueue-1;i>=0;i--{
				result = append(result, queue[i].Val)
				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
				}
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
				}
			}
			reverse=false
		}else {
			for i:=lenQueue-1;i>=0;i--{
				result = append(result, queue[i].Val)
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
				}
				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
				}
			}
			reverse=true
		}
		results = append(results, result)
		queue=queue[lenQueue:]
	}

	return results
}
