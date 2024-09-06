package problem389

import (
	"fmt"
	"testing"
)

func Test_389(t *testing.T) {
	fmt.Println(findTheDifference("afdfsdfdsfs", "aafdfsdfdsfs"))
}

func findTheDifference(s string, t string) byte {
	sCount:=make(map[rune]int)
	tCount:=make(map[rune]int)

	for _,v:=range s{
		sCount[v]++
	}
	for _,v:=range t{
		tCount[v]++
	}
	for k,v:=range tCount{
		if v!=sCount[k]{
			return byte(k)
		}
	}
	return '0'
}
