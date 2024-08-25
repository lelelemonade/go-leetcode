package problem145

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversalRecursion(root *TreeNode)(res []int){
	if root==nil{
		return []int{}
	}
	left:=postorderTraversalRecursion(root.Left)
	right:=postorderTraversalRecursion(root.Right)
	return append(append(left,right...),root.Val)
}

func postorderTraversal(root *TreeNode) (res []int) {
	var stack []*TreeNode
	for root != nil {
		if root.Left==nil&&root.Right==nil{
			res = append(res, root.Val)
		}else if root.Left!=nil&&root.Right==nil{
			stack = append(stack, root)
			root=root.Left
		}else if root.Left==nil&&root.Right!=nil{
			stack = append(stack, root)
			root=root.Right
		}else{
			stack = append(stack, root)
			root=root.Left
		}
	}
	return
}
