package problem117

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
	queue := make([]*Node, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		queueLen := len(queue)

		for i := 0; i < queueLen; i++ {
			if i != queueLen-1 {
				queue[i].Next = queue[i+1]
			}

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[queueLen:]
	}

	return root
}
