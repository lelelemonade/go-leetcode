package problem116

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	stack := make([]*Node, 0)
	stack = append(stack, root)

	for len(stack) != 0 {
		levelLength := len(stack)

		for i := 0; i < levelLength; i++ {
			currentNode := stack[0]
			if i < levelLength-1 {
				currentNode.Next = stack[1]
			}
			stack = stack[1:]
			if currentNode.Left != nil {
				stack = append(stack, currentNode.Left)
			}
			if currentNode.Right != nil {
				stack = append(stack, currentNode.Right)
			}
		}
	}

	return root
}
