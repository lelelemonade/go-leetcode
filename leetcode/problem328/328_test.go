package problem328

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	dummyOdd:=&ListNode{}
	dummyEven:=&ListNode{}

	oddIter:=dummyOdd
	evenIter:=dummyEven
	headIter:=head
	count:=0

	for headIter!=nil{
		count++
		if count%2==0{
			evenIter.Next=headIter
			evenIter=evenIter.Next
			
		}else{
			oddIter.Next=headIter
			oddIter=oddIter.Next
			
		}
		headIter=headIter.Next
		evenIter.Next=nil
		oddIter.Next=nil
	}

	oddIter.Next=dummyEven.Next

	return dummyOdd.Next
}

func Test328(t *testing.T){
	node1:=&ListNode{1,nil}
	node2:=&ListNode{2,node1}
	node3:=&ListNode{3,node2}
	node4:=&ListNode{4,node3}
	node5:=&ListNode{5,node4}
	fmt.Println(oddEvenList(node5))
}
