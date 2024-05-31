package problem501

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}

	currentMaxVals := make(map[int]struct{})
	//currentMaxVals[root.Val] = struct{}{}
	currenMaxCount := 0

	dfs(root, &currentMaxVals, &currenMaxCount)

	result := make([]int, 0)
	for k, _ := range currentMaxVals {
		result = append(result, k)
	}

	return result
}

func dfs(root *TreeNode, currentMaxVals *map[int]struct{}, currenMaxCount *int) {
	if root == nil {
		return
	}

	if _, e := (*currentMaxVals)[root.Val]; !e {
		searchModeResult := 0
		searchMode(root, root.Val, &searchModeResult)
		if *currenMaxCount == searchModeResult {
			(*currentMaxVals)[root.Val] = struct{}{}
		} else if *currenMaxCount < searchModeResult {
			*currentMaxVals = make(map[int]struct{})
			(*currentMaxVals)[root.Val] = struct{}{}
			*currenMaxCount = searchModeResult
		}
	}

	dfs(root.Left, currentMaxVals, currenMaxCount)
	dfs(root.Right, currentMaxVals, currenMaxCount)
}

func searchMode(root *TreeNode, currentMaxVal int, currenMaxCount *int) {
	if root == nil || root.Val != currentMaxVal {
		return
	}
	*currenMaxCount++
	searchMode(root.Left, currentMaxVal, currenMaxCount)
	searchMode(root.Right, currentMaxVal, currenMaxCount)
}
