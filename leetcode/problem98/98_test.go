package problem98

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	inOrderArray := dfs(root)

	for i := 1; i < len(inOrderArray); i++ {
		if inOrderArray[i] <= inOrderArray[i-1] {
			return false
		}
	}
	return true
}

func dfs(root *TreeNode) (result []int) {
	if root == nil {
		return nil
	}

	leftResult := dfs(root.Left)
	if leftResult != nil {
		result = append(leftResult, root.Val)
	} else {
		result = []int{root.Val}
	}

	rightResult := dfs(root.Right)

	if rightResult != nil {
		result = append(result, rightResult...)
	}

	return result
}
