package problem1161

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	queue:=make([]*TreeNode,1)
	queue[0]=root
	maxSum:=-100001
	layerNumber:=0
	result:=0

	for len(queue)!=0{
		lenQueue:=len(queue)
		sum:=0

		for i:=0;i<lenQueue;i++{
			sum+=queue[i].Val
			if queue[i].Left!=nil{
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right!=nil{
				queue = append(queue, queue[i].Right)
			}
		}
		queue=queue[lenQueue:]
		layerNumber++
		if sum>maxSum{
			maxSum=sum
			result=layerNumber
		}
	}

	return result
}
