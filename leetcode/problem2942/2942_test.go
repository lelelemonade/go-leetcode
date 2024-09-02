package problem2942

import "strings"

func findWordsContaining(words []string, x byte) []int {
    result:=make([]int,0)

	for i,v:=range words{
		if strings.Contains(v,string(x)){
			result=append(result, i)
		}
	}
	return result
}