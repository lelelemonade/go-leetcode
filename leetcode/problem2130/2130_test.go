package problem2130

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	iter := head
	nodeCount := 0
	for iter != nil {
		iter = iter.Next
		nodeCount++
	}
	iter = head
	result := 0
	valueMap := make(map[int]int)
	for i := 0; i < nodeCount/2; i++ {
		valueMap[i] = iter.Val
		iter = iter.Next
	}
	for i := nodeCount / 2; i < nodeCount; i++ {
		result = max(result, iter.Val+valueMap[nodeCount-1-i])
		iter = iter.Next
	}

	return result
}
