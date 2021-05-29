package problem203

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	sentinel := &ListNode{Next: head}
	before := sentinel
	curr := head
	for {
		if curr == nil {
			break
		}
		if curr.Val == val {
			before.Next = curr.Next
		} else {
			before = before.Next
		}
		curr = curr.Next
	}
	return sentinel.Next
}
