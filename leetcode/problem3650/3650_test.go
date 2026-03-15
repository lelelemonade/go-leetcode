package problem3650

import (
	"container/heap"
	"math"
)

type MinHeap [][2]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
    *h = append(*h, x.([2]int))
}

func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func minCost(n int, edges [][]int) int {
    graph := make([][][2]int, n) // graph[u] = [[v, cost], ...]
    
    for _, edge := range edges {
        u, v, w := edge[0], edge[1], edge[2]
        graph[u] = append(graph[u], [2]int{v, w})
        graph[v] = append(graph[v], [2]int{u, 2 * w})
    }
    
    dist := make([]int, n)
    for i := range dist {
        dist[i] = math.MaxInt
    }
    dist[0] = 0
    
    pq := &MinHeap{}
    heap.Init(pq)
    heap.Push(pq, [2]int{0, 0})
    
    for pq.Len() > 0 {
        curr := heap.Pop(pq).([2]int)
        d, u := curr[0], curr[1]
        
        if d > dist[u] {
            continue
        }
        
        for _, neighbor := range graph[u] {
            v, w := neighbor[0], neighbor[1]
            newDist := d + w
            
            if newDist < dist[v] {
                dist[v] = newDist
                heap.Push(pq, [2]int{newDist, v})
            }
        }
    }
    
    if dist[n-1] == int(1e9) {
        return -1
    }
    return dist[n-1]
}