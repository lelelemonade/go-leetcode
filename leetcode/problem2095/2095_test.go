package problem2095

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next==nil{
		return nil
	}

	slow:=&ListNode{0,head}
	fast:=head

	for fast!=nil&&fast.Next!=nil{
		fast=fast.Next.Next
		slow=slow.Next
	}

	slow.Next=slow.Next.Next

	return head
}
