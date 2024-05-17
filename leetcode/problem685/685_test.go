package problem685

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test685(t *testing.T) {
	//testCase1 := [][]int{{1, 2}, {1, 3}, {2, 3}}
	//
	//assert.Equal(t, []int{2, 3}, findRedundantDirectedConnection(testCase1))
	//
	//testCase2 := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}}
	//
	//assert.Equal(t, []int{4, 1}, findRedundantDirectedConnection(testCase2))
	//
	//testCase3 := [][]int{{2, 1}, {3, 1}, {4, 2}, {1, 4}}
	//
	//assert.Equal(t, []int{2, 1}, findRedundantDirectedConnection(testCase3))

	//testCase4 := [][]int{{3, 2}, {2, 1}, {1, 3}, {4, 1}}
	testCase4 := [][]int{{3, 2}, {1, 3}, {2, 1}, {4, 1}}
	assert.Equal(t, []int{2, 1}, findRedundantDirectedConnection(testCase4))
}

func findRedundantDirectedConnection(edges [][]int) (redundantEdge []int) {
	n := len(edges)
	uf := newUnionFind(n + 1)
	parent := make([]int, n+1) // parent[i] 表示 i 的父节点
	for i := range parent {
		parent[i] = i
	}

	var conflictEdge, cycleEdge []int
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		if parent[to] != to { // to 有两个父节点
			conflictEdge = edge
		} else {
			parent[to] = from
			if uf.findAncestor(from) == uf.findAncestor(to) { // from 和 to 已连接
				cycleEdge = edge
			} else {
				uf.union(from, to)
			}
		}
	}

	// 若不存在一个节点有两个父节点的情况，则附加的边一定导致环路出现
	if conflictEdge == nil {
		return cycleEdge
	}
	// conflictEdge[1] 有两个父节点，其中之一与其构成附加的边
	// 由于我们是按照 edges 的顺序连接的，若在访问到 conflictEdge 之前已经形成了环路，则附加的边在环上
	// 否则附加的边就是 conflictEdge
	if cycleEdge != nil {
		return []int{parent[conflictEdge[1]], conflictEdge[1]}
	}
	return conflictEdge
}

type unionFind struct {
	ancestor []int
}

func newUnionFind(n int) unionFind {
	ancestor := make([]int, n)
	for i := 0; i < n; i++ {
		ancestor[i] = i
	}
	return unionFind{ancestor}
}

func (uf unionFind) findAncestor(x int) int {
	if uf.ancestor[x] != x {
		uf.ancestor[x] = uf.findAncestor(uf.ancestor[x])
	}
	return uf.ancestor[x]
}

func (uf unionFind) union(from, to int) {
	uf.ancestor[uf.findAncestor(from)] = uf.findAncestor(to)
}

//func findRedundantDirectedConnection(edges [][]int) []int {
//	unionFind := NewUnionFind(len(edges))
//
//	for i := 0; i < len(edges); i++ {
//		unionFind.directionUnion(edges[i][0], edges[i][1])
//	}
//
//	//1. no 2 input point
//	if unionFind.overrideParentIndex == 0 {
//		iter := unionFind.circularDetect(1)
//		circularMap := unionFind.getCircularPoints(iter)
//
//		for i := len(edges) - 1; i > 0; i-- {
//			_, startInCircle := circularMap[edges[i][0]]
//			_, endInCircle := circularMap[edges[i][1]]
//			if startInCircle && endInCircle {
//				return edges[i]
//			}
//		}
//
//		return []int{0, 0}
//	}
//
//	//2. two input point and circular
//	if unionFind.circularDetect(unionFind.overrideParent) != 0 {
//		return []int{unionFind.overrideParentIndex, unionFind.overrideParent}
//	} else if unionFind.circularDetect(unionFind.overrideParentIndex) != 0 {
//		return []int{unionFind.overrideParentIndex, unionFind.parent[unionFind.overrideParentIndex]}
//	}
//
//	//3. two input point and no circular
//	for i := len(edges) - 1; i >= 0; i-- {
//		if edges[i][1] == unionFind.overrideParentIndex && edges[i][0] == unionFind.overrideParent {
//			return edges[i]
//		}
//
//		if edges[i][1] == unionFind.overrideParentIndex && unionFind.parent[edges[i][0]] == unionFind.parent[unionFind.overrideParentIndex] {
//			return edges[i]
//		}
//	}
//
//	return []int{}
//}
//
//type UnionFind struct {
//	ancestor              []int
//	parent              []int
//	overrideParent      int
//	overrideParentIndex int
//}
//
//func NewUnionFind(n int) UnionFind {
//	parent := make([]int, n+1)
//	ancestor := make([]int, n+1)
//
//	for i := 0; i < n+1; i++ {
//		parent[i] = i
//		ancestor[i] = i
//	}
//
//	return UnionFind{ancestor,parent, 0, 0}
//}
//
//func (u *UnionFind) circularDetect(start int) int {
//	visited := make(map[int]struct{})
//	_, exist := visited[start]
//
//	for !exist {
//		visited[start] = struct{}{}
//		start = u.parent[start]
//		_, exist = visited[start]
//	}
//
//	if start == u.parent[start] {
//		return 0
//	}
//
//	return start
//}
//
//func (u *UnionFind) getCircularPoints(start int) map[int]struct{} {
//	visited := make(map[int]struct{})
//	_, exist := visited[start]
//
//	for !exist {
//		visited[start] = struct{}{}
//		start = u.parent[start]
//		_, exist = visited[start]
//	}
//
//	return visited
//}
//
//func (u *UnionFind) findAncestor(x int) int {
//	ancestor := u.ancestor[x]
//
//	if ancestor == x {
//		return ancestor
//	}
//	return u.findAncestor(ancestor)
//}
//
//func (u *UnionFind) directionUnion(x int, y int) {
//	xAncestor := u.findAncestor(x)
//	yAncestor := u.findAncestor(y)
//
//	if y != yAncestor {
//		u.overrideParent = yAncestor
//		u.overrideParentIndex = y
//	}
//
//	u.ancestor[yAncestor] = xAncestor
//	u.parent[y] = x
//}
