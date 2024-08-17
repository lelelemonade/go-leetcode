package problem141

type ListNode struct {
    Val int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
	dummy:=&ListNode{Next: head}
    fast:=head
	slow:=dummy

	for fast != slow{
		if fast==nil||fast.Next==nil{
			return false
		}
		fast=fast.Next.Next
		slow=slow.Next
	}
	return true
}