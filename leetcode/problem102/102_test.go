package problem102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return make([][]int, 0)
	}

	result := make([][]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) != 0 {
		levelResult := make([]int, 0)
		levelLength := len(stack)
		for i := 0; i < levelLength; i++ {
			currentNode := stack[0]
			stack = stack[1:]
			if currentNode.Left != nil {
				stack = append(stack, currentNode.Left)
			}
			if currentNode.Right != nil {
				stack = append(stack, currentNode.Right)
			}
			levelResult = append(levelResult, currentNode.Val)

		}
		result = append(result, levelResult)
	}

	return result
}
