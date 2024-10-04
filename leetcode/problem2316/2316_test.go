package problem2316

type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{make([]int, n), make([]int, n)}

	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX != rootY {
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX]++
		}
	}
}

func countPairs(n int, edges [][]int) int64 {
	unionFind := NewUnionFind(n)
	for _, edge := range edges {
		unionFind.Union(edge[0], edge[1])
	}
	sizes := make(map[int]int)
	for i := 0; i < n; i++ {
		sizes[unionFind.Find(i)]++
	}

	result := 0
	for i := 0; i < n; i++ {
		result += n - sizes[unionFind.Find(i)]
	}

	return int64(result)
}
