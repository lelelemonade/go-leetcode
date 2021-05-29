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
