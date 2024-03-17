package problem654

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	var maxValue int
	var maxIndex int

	for i, v := range nums {
		if v > maxValue {
			maxValue = v
			maxIndex = i
		}
	}
	return &TreeNode{
		Val:   maxValue,
		Left:  constructMaximumBinaryTree(nums[:maxIndex]),
		Right: constructMaximumBinaryTree(nums[maxIndex+1:]),
	}

}
