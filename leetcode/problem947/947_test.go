package problem947

func removeStones(stones [][]int) int {
	uf := make(map[int]int)
	unionCount := 0

	var find func(x int) int
	find = func(x int) int {
		if _, e := uf[x]; !e {
			uf[x] = x
			unionCount++
		}
		if x != uf[x] {
			uf[x] = find(uf[x])
		}
		return uf[x]
	}

	union := func(x, y int) {
		rootX := find(x)
		rootY := find(y)
		if rootX == rootY {
			return
		}
		uf[rootX] = rootY
		unionCount--
	}

	for _, v := range stones {
		union(v[0]+10001, v[1])
	}
	return len(stones) - unionCount
}
