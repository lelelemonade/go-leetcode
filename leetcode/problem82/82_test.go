package problem82

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	iter := dummy

	for iter.Next != nil && iter.Next.Next!=nil {
		if iter.Next.Val == iter.Next.Next.Val {
			valToDel:=iter.Next.Val
			for iter.Next!=nil && iter.Next.Val == valToDel{
				iter.Next=iter.Next.Next
			}
		}else{
			iter=iter.Next
		}
	}
	return dummy.Next
}
