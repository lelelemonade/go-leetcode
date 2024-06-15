package problem206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	iter := head
	var prev *ListNode

	for iter != nil {
		temp := iter.Next
		iter.Next = prev
		prev = iter
		iter = temp
	}

	return prev
}
