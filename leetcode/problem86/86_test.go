package problem86

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	dummySmall:=&ListNode{}
	small:=dummySmall
	dummyBig:=&ListNode{}
	big:=dummyBig

	for head!=nil {
		if head.Val<x{
			small.Next=head
			small=small.Next
		}else{
			big.Next=head
			big=big.Next
		}
		head=head.Next
	}

	big.Next=nil
	small.Next=dummyBig.Next

	return dummySmall.Next
}

func Test86(t *testing.T) {
	node1:=&ListNode{Val:2,Next: nil}
	node2:=&ListNode{Val:5,Next: node1}
	node3:=&ListNode{Val:2,Next: node2}
	node4:=&ListNode{Val:3,Next: node3}
	node5:=&ListNode{Val:4,Next: node4}
	node6:=&ListNode{Val:1,Next: node5}
	fmt.Println(partition(node6,3))
}