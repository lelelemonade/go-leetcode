package problem1019

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	stack := make([]*ListNode, 0)
	valuesMap := make(map[*ListNode]int)
	result := make([]int, 0)

	iter := head
	i := 0
	for iter != nil {
		valuesMap[iter] = i

		for len(stack) > 0 && stack[len(stack)-1].Val < iter.Val {
			result[valuesMap[stack[len(stack)-1]]] = iter.Val
			stack=stack[:len(stack)-1]
		}
		stack = append(stack, iter)
		result = append(result, 0)
		iter=iter.Next
		i++
	}
	return result
}
