package problem24

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}
	iter := dummy

	for iter != nil && iter.Next != nil && iter.Next.Next != nil {
		nextIter := iter.Next.Next.Next

		iter.Next.Next.Next = iter.Next
		iter.Next = iter.Next.Next
		iter.Next.Next.Next = nextIter

		iter = iter.Next.Next
	}

	return dummy.Next
}
