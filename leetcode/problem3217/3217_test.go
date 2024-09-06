package problem3217

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	numsMap:=make(map[int]struct{})
	for _,v:=range nums{
		numsMap[v]=struct{}{}
	}

	dummy:=&ListNode{Next:head}
	iter:=dummy

	for iter.Next!=nil{
		if _,e:=numsMap[iter.Next.Val];e{
			iter.Next=iter.Next.Next
		}else{
			iter=iter.Next
		}
	}
	return dummy.Next
}
