package problem173

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	root *TreeNode
	iter *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{root: root,iter: root}
}

func (this *BSTIterator) Next() int {

}

func (this *BSTIterator) HasNext() bool {

}
