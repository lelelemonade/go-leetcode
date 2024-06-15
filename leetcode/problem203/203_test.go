package problem203

import (
	"fmt"
	"testing"
)

func Test_Problem203(t *testing.T) {
	testArray := []int{1, 2, 6, 3, 4, 5, 6}

	head := &ListNode{Val: testArray[0]}
	temp := head

	for _, val := range testArray {
		temp.Next = &ListNode{Val: val}
		temp = temp.Next
	}

	fmt.Println(removeElements(head, 6))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	iter := dummy

	for iter.Next != nil {
		if iter.Next.Val == val {
			iter.Next = iter.Next.Next
		} else {
			iter = iter.Next
		}
	}

	return dummy.Next
}
