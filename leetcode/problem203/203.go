package problem203

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) (newHead *ListNode) {
	newHead = head
	before := &ListNode{Next: head}
	for head.Next != nil {
		if head.Val == val {
			before.Next = head.Next
		} else {
			before = before.Next
		}
		head = head.Next
	}

	return newHead
}
