package problem143

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	for head!=nil&& head.Next!=nil{
		secondToLast:=head
		for secondToLast.Next.Next!=nil{
			secondToLast=secondToLast.Next
		}
		last:=secondToLast.Next
		secondToLast.Next=nil
		last.Next=head.Next
		head.Next=last
		head=last.Next
	}
}
