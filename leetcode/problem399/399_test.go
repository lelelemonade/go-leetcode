package problem399

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    equationMap:=make(map[string]map[string]float64)

	for i:=0;i<len(equations);i++{
		if equationMap[equations[i][0]]==nil{
			equationMap[equations[i][0]]=make(map[string]float64)
		}
		if equationMap[equations[i][1]]==nil{
			equationMap[equations[i][1]]=make(map[string]float64)
		}
		equationMap[equations[i][0]][equations[i][1]]=values[i]
		equationMap[equations[i][1]][equations[i][0]]=1/values[i]
	}

	visited:=make(map[string]struct{})
	var dfs func(start string,end string)float64
	dfs=func(start string,end string)float64{
		defer delete(visited,start)
		if equationMap[start][end]!=0{
			return equationMap[start][end]
		}
		for k,v:=range equationMap[start]{
			if _,e:=visited[k];e{
				continue
			}
			visited[k]=struct{}{}
			result:=dfs(k,end)
			if result !=-1{
				return v*result
			}
		}
		return -1
	}

	result:=make([]float64, 0)
	for _,v:=range queries{
		result = append(result, dfs(v[0],v[1]))
	}
	return result
}

func Test399(t *testing.T) {
	assert.Equal(t,[]float64{6.00000,0.50000,-1.00000,1.00000,-1.00000},calcEquation(
		[][]string{{"a","b"},{"b","c"}},
		[]float64{2.0,3.0},
		[][]string{{"a","c"},{"b","a"},{"a","e"},{"a","a"},{"x","x"}},
	))
}