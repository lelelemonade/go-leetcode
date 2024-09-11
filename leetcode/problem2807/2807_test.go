package problem2807

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	var getgcd func(a,b int)int
	getgcd=func(a,b int)int{
		if a%b==0{
			return b
		}
		return getgcd(b,a%b)
	}
	iter:=head
	for iter.Next!=nil{
		gcdNode:=&ListNode{Val: getgcd(iter.Val,iter.Next.Val)}
		gcdNode.Next=iter.Next
		iter.Next=gcdNode
		iter=gcdNode.Next
	}
	return head
}
