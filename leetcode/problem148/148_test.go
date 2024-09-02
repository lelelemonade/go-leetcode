package problem148

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}
func merge(head1, head2 *ListNode) *ListNode {
    dummyHead := &ListNode{}
    temp, temp1, temp2 := dummyHead, head1, head2
    for temp1 != nil && temp2 != nil {
        if temp1.Val <= temp2.Val {
            temp.Next = temp1
            temp1 = temp1.Next
        } else {
            temp.Next = temp2
            temp2 = temp2.Next
        }
        temp = temp.Next
    }
    if temp1 != nil {
        temp.Next = temp1
    } else if temp2 != nil {
        temp.Next = temp2
    }
    return dummyHead.Next
}

func sortList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }

    length := 0
    for node := head; node != nil; node = node.Next {
        length++
    }

    dummyHead := &ListNode{Next: head}
    for subLength := 1; subLength < length; subLength <<= 1 {
        prev, cur := dummyHead, dummyHead.Next
        for cur != nil {
            head1 := cur
            for i := 1; i < subLength && cur.Next != nil; i++ {
                cur = cur.Next
            }

            head2 := cur.Next
            cur.Next = nil
            cur = head2
            for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
                cur = cur.Next
            }

            var next *ListNode
            if cur != nil {
                next = cur.Next
                cur.Next = nil
            }

            prev.Next = merge(head1, head2)

            for prev.Next != nil {
                prev = prev.Next
            }
            cur = next
        }
    }
    return dummyHead.Next
}

func Test148(t *testing.T){
	node4:=&ListNode{Val:3,Next: nil}
	node3:=&ListNode{Val:1,Next: node4}
	node2:=&ListNode{Val:2,Next: node3}
	node1:=&ListNode{Val:4,Next: node2}

	result:=sortList(node1)
	assert.Equal(t,node1,result)
}
