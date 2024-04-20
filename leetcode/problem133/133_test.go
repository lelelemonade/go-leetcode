package problem133

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	visited := make(map[*Node]*Node)
	return dfs(node,visited)
}

func dfs(node *Node,visted map[*Node]*Node) *Node {
	if node == nil {
		return node
	}
	if v,exist:=visted[node];exist{
		return v
	}else{
		newNode := &Node{node.Val, make([]*Node, len(node.Neighbors))}

		visted[node]=newNode

		for i := 0; i < len(node.Neighbors); i++ {
			newNode.Neighbors[i] = dfs(node.Neighbors[i],visted)
		}

		return newNode
	}
}

func Test133(t *testing.T) {
	node1 := &Node{1, nil}
	node2 := &Node{2, nil}
	node3 := &Node{3, nil}
	node4 := &Node{4, nil}
	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}

	newNode := cloneGraph(node1)

	assert.NotEqual(t, newNode, node1)
}
