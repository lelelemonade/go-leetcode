package problem2933

import (
	"slices"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func findHighAccessEmployees(access_times [][]string) []string {
    result:=make([]string,0)
	accessMap:=make(map[string][]int)

	for _,v:=range access_times{
		if _,e:=accessMap[v[0]];!e{
			accessMap[v[0]]=make([]int, 0)
		}
		hour,e:=strconv.Atoi(v[1][:2])
		if e!=nil{
			panic(e)
		}
		minute,e:=strconv.Atoi(v[1][2:])
		if e!=nil{
			panic(e)
		}
		accessMap[v[0]]=append(accessMap[v[0]], hour*60+minute)
	}

	for k,v:=range accessMap{
		if len(v)<3{
			continue
		}
		slices.Sort(v)

		for i:=0;i<len(v)-2;i++{
			if v[i+2]-v[i]<60{
				result = append(result, k)
				break
			}
		}
	}

	return result
}

func Test2933(t *testing.T) {
	assert.Equal(t,[]string{"dz"},findHighAccessEmployees([][]string{{"dz","0719"},{"dz","0726"},{"dz","0716"},{"dz","0716"}}))
}