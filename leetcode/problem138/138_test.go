package problem138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	oldNodeToIdxMap := make(map[*Node]int)
	newIdxToNodeMap := make(map[int]*Node)

	newDummyHead := &Node{}
	newIter := newDummyHead

	oldIter := head
	idx := 0
	for oldIter != nil {
		newTemp := &Node{
			Val: oldIter.Val,
		}
		newIter.Next = newTemp
		newIter = newTemp

		oldNodeToIdxMap[oldIter] = idx
		newIdxToNodeMap[idx] = newIter
		oldIter = oldIter.Next
        idx++
	}

	oldIter = head
	newIter = newDummyHead.Next
	for oldIter != nil {
        if oldIter.Random != nil {
            newIter.Random = newIdxToNodeMap[oldNodeToIdxMap[oldIter.Random]]
        }
		oldIter = oldIter.Next
        newIter=newIter.Next
	}

	return newDummyHead.Next
}
