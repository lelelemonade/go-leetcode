package problem92

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	iter := dummy
	var last *ListNode
	var beforeRangeNode *ListNode
	var rangeStartNode *ListNode
	idx := 0

	for iter != nil {
		if idx == left-1{
			beforeRangeNode=iter
		}else if idx == left{
			rangeStartNode=iter
		}

		if idx >= left && idx <= right {
			temp:=iter.Next
			iter.Next=last
			last=iter
			if idx == right{
				beforeRangeNode.Next=iter
				rangeStartNode.Next=temp
			}
			iter=temp
		}else{
			last=iter
			iter=iter.Next
		}
		idx++
	}
	return dummy.Next
}

func Test92(t *testing.T) {
	testNode1:=&ListNode{Val:5,Next: nil}
	testNode2:=&ListNode{Val:4,Next: testNode1}
	testNode3:=&ListNode{Val:3,Next: testNode2}
	testNode4:=&ListNode{Val:2,Next: testNode3}
	testNode5:=&ListNode{Val:1,Next: testNode4}

	fmt.Println(reverseBetween(testNode5,2,4))
}