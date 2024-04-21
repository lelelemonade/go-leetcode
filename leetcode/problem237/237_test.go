package problem237

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	for iter := node; iter.Next != nil; iter = iter.Next {
		iter.Val = iter.Next.Val
		if iter.Next.Next == nil {
			iter.Next = nil
			break
		}
	}
}
