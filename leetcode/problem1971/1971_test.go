package problem1971

func validPath(n int, edges [][]int, source int, destination int) bool {
	unionFind := NewUnionFind(n)

	for i := 0; i < len(edges); i++ {
		unionFind.merge(edges[i][0], edges[i][1])
	}

	return unionFind.find(source) == unionFind.find(destination)
}

type UnionFind struct {
	f    []int
	rank []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		f:    make([]int, n),
		rank: make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.f[i] = i
	}
	return uf
}

func (uf *UnionFind) merge(x, y int) {
	rx := uf.find(x)
	ry := uf.find(y)
	if rx != ry {
		if uf.rank[rx] > uf.rank[ry] {
			uf.f[ry] = rx
		} else if uf.rank[rx] < uf.rank[ry] {
			uf.f[rx] = ry
		} else {
			uf.f[ry] = rx
			uf.rank[rx]++
		}
	}
}

func (uf *UnionFind) find(x int) int {
	if x != uf.f[x] {
		uf.f[x] = uf.find(uf.f[x])
	}
	return uf.f[x]
}
