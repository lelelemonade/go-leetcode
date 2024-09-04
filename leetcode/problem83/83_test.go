package problem83

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	iter:=head

	for iter!=nil&&iter.Next!=nil{
		if iter.Next.Val==iter.Val{
			iter.Next=iter.Next.Next
		}else{
			iter=iter.Next
		}
	}
	return head
}
