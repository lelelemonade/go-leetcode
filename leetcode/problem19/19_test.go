package problem19

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}

	fast := dummy
	slow := dummy
	i := 0

	for fast != nil {
		fast = fast.Next
		if i > n {
			slow = slow.Next
		}
		i++
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
