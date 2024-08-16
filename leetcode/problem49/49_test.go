package problem49

import (
	"fmt"
	"testing"
)

func groupAnagrams(strs []string) [][]string {
	countToAnagrams:=make(map[[26]int][]string,0)
	for i:=0;i<len(strs);i++{
		count:=[26]int{}
		for j:=0;j<len(strs[i]);j++{
			count[strs[i][j]-'a']++
		}
		countToAnagrams[count]=append(countToAnagrams[count], strs[i])
	}
	results:=make([][]string,0)
	for _,v:=range countToAnagrams{
		results=append(results, v)
	}
	return results
}

func Test49(t *testing.T) {
	fmt.Print(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
