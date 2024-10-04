package problem797

func allPathsSourceTarget(graph [][]int) [][]int {
	results:=make([][]int,0)
    var dfs func(iterated []int)
	dfs = func(iterated []int){
		x:=iterated[len(iterated)-1]
		if x==len(graph)-1{
			result:=make([]int,len(iterated))
			copy(result,iterated)
			results = append(results, result)
			return
		}
		for _,v:=range graph[x]{
			dfs(append(iterated,v))
		}
	}
	dfs([]int{0})
	return results
}