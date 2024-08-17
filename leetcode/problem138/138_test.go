package problem138

type Node struct {
    Val int
    Next *Node
    Random *Node
}


 func copyRandomList(head *Node) *Node {
    nextMap:=make(map[*Node]*Node)
	nextMap:=make(map[*Node]*Node)
}