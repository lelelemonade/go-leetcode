package problem2816

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Test2816(t *testing.T) {
	//testCase1ListNode1 := &ListNode{9, nil}
	//testCase1ListNode2 := &ListNode{9, testCase1ListNode1}
	//testCase1ListNode3 := &ListNode{9, testCase1ListNode2}
	//
	//printNode := doubleIt(testCase1ListNode3)
	//for printNode != nil {
	//	fmt.Print(printNode.Val)
	//	printNode = printNode.Next
	//}
	//
	//testCase2ListNode1 := &ListNode{9, nil}
	//testCase2ListNode2 := &ListNode{8, testCase2ListNode1}
	//testCase2ListNode3 := &ListNode{1, testCase2ListNode2}
	//
	//printNode = doubleIt(testCase2ListNode3)
	//for printNode != nil {
	//	fmt.Print(printNode.Val)
	//	printNode = printNode.Next
	//}

	testCase3ListNode1 := &ListNode{3, nil}

	printNode := doubleIt(testCase3ListNode1)
	for printNode != nil {
		fmt.Print(printNode.Val)
		printNode = printNode.Next
	}
}

func doubleIt(head *ListNode) *ListNode {
	nodeStack := make([]*ListNode, 0)

	iter := head

	for iter != nil {
		nodeStack = append(nodeStack, iter)
		iter = iter.Next
	}

	carry := 0

	for i := len(nodeStack) - 1; i >= 0; i-- {
		multiply := 2*nodeStack[i].Val + carry
		nodeStack[i].Val = multiply % 10
		carry = multiply / 10

	}

	if carry != 0 {
		return &ListNode{Val: carry, Next: head}
	} else {
		return head
	}
}
