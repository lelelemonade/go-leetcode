package problem35

import (
	"fmt"
	"testing"
)

func Test_Problem35(t *testing.T) {
	testArray := []int{1, 3, 5, 6}

	fmt.Println(searchInsert(testArray, 5))
	fmt.Println(testArray)
}
