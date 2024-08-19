package problem61

import (
	"fmt"
	"testing"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	iter:=head
	length:=0
	for iter!=nil {
		length++
		iter=iter.Next
	}

	if length == 0 {
		return head
	}

	moveStep := k%length
	fast:=head
	slow:=head
	idx:=0
	for fast.Next!=nil{
		fast=fast.Next
		idx++
		if idx>moveStep{
			slow=slow.Next
		}
	}
	fast.Next=head
	result:=slow.Next
	slow.Next=nil
	return result
}

func Test61(t *testing.T) {
	node1:=&ListNode{Val:5,Next: nil}
	node2:=&ListNode{Val:4,Next: node1}
	node3:=&ListNode{Val:3,Next: node2}
	node4:=&ListNode{Val:2,Next: node3}
	node5:=&ListNode{Val:1,Next: node4}
	fmt.Println(rotateRight(node5,2))
}