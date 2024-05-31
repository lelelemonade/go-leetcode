package problem199

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}

	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) != 0 {
		levelLength := len(stack)

		currentNode := stack[0]
		for i := 0; i < levelLength; i++ {
			currentNode = stack[0]
			stack = stack[1:]
			if currentNode.Left != nil {
				stack = append(stack, currentNode.Left)
			}
			if currentNode.Right != nil {
				stack = append(stack, currentNode.Right)
			}
		}
		result = append(result, currentNode.Val)
	}

	return result
}
