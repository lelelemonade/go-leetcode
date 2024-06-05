package problem2196

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodeMap := make(map[int]*TreeNode)
	childMap := make(map[int]struct{})

	for _, node := range descriptions {
		treeNode := nodeMap[node[0]]
		if treeNode == nil {
			treeNode = &TreeNode{
				Val:   node[0],
				Left:  nil,
				Right: nil,
			}
			nodeMap[node[0]] = treeNode
		}
		childNode := nodeMap[node[1]]
		if childNode == nil {
			childNode = &TreeNode{
				Val:   node[1],
				Left:  nil,
				Right: nil,
			}
			nodeMap[node[1]] = childNode
		}

		if node[2] == 1 {
			treeNode.Left = childNode
		} else {
			treeNode.Right = childNode
		}
		childMap[node[1]] = struct{}{}
	}
	for key, _ := range nodeMap {
		if _, exist := childMap[key]; !exist {
			return nodeMap[key]
		}
	}
	return nil
}
