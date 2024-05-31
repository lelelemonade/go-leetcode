package problem429

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return make([][]int, 0)
	}

	result := make([][]int, 0)
	stack := make([]*Node, 0)
	stack = append(stack, root)

	for len(stack) != 0 {
		levelLength := len(stack)

		levelResult := make([]int, 0)
		for i := 0; i < levelLength; i++ {
			currentNode := stack[0]
			stack = stack[1:]

			for _, v := range currentNode.Children {
				stack = append(stack, v)
			}
			levelResult = append(levelResult, currentNode.Val)
		}
		result = append(result, levelResult)
	}

	return result
}
