package problem21

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	iter := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			iter.Next = list1
			iter = iter.Next
			list1 = list1.Next
		} else {
			iter.Next = list2
			iter = iter.Next
			list2 = list2.Next
		}
	}

	if list1 == nil {
		iter.Next = list2
	} else if list2 == nil {
		iter.Next = list1
	}

	return dummy.Next
}
