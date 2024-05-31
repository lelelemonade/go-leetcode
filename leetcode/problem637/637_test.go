package problem637

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return make([]float64, 0)
	}

	result := make([]float64, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) != 0 {
		levelLength := len(stack)

		total := 0
		for i := 0; i < levelLength; i++ {
			currentNode := stack[0]
			stack = stack[1:]
			if currentNode.Left != nil {
				stack = append(stack, currentNode.Left)
			}
			if currentNode.Right != nil {
				stack = append(stack, currentNode.Right)
			}
			total += currentNode.Val
		}
		result = append(result, (float64(total))/float64(levelLength))
	}

	return result
}
