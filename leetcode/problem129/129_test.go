package problem129

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	digits:=make([]int,0)

	total:=0

	sumDFS(root,digits,&total)

	return total
}

func sumDFS(root *TreeNode,digits []int,total *int){
	if root == nil {
		return
	}
	if root.Left==nil && root.Right==nil {
		*total+=sumDigitArray(append(digits,root.Val))
	}

	sumDFS(root.Left,append(digits,root.Val),total)
	sumDFS(root.Right,append(digits,root.Val),total)
}

func sumDigitArray(digits []int)int {
	result:=0
	j:=1

	for i:=len(digits)-1;i>=0;i--{
		result+=j*digits[i]
		j*=10
	}
	return result
}