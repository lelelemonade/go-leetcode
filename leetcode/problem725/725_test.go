package problem725

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	size:=0
	iter:=&ListNode{Next: head}
	for iter.Next!=nil{
		iter=iter.Next
		size++
	}
	iter=head
	result:=make([]*ListNode,0)

	addPart:=func(listCount int,listSize int){
		for i:=0;i<listCount;i++{
			if listSize<0{
				result = append(result, nil)
				continue
			}
			part:=iter
			for j:=0;j<listSize;j++{
				iter=iter.Next
			}
			temp:=iter.Next
			iter.Next=nil
			iter=temp
			result = append(result, part)
		}
	}

	addPart(size%k,size/k)
	addPart(k-size%k,size/k-1)

	return result
}