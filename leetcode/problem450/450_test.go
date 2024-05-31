package problem450

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestSlices(t *testing.T) {
	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := sliceA[:3]
	sliceB[0] = 6
	fmt.Println(sliceA)
}
