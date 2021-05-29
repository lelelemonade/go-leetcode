package problem203

import (
	"fmt"
	"testing"
)

func Test_Problem203(t *testing.T) {
	testArray := []int{6, 6, 6, 6}

	head := ListNode{Val: testArray[0]}
	temp := head

	for _, val := range testArray {
		temp.Next = &ListNode{Val: val}
		temp = *temp.Next
	}

	fmt.Println(removeElements(&head, 6))
}
