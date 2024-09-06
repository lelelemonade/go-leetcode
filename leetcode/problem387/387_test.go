package problem387

func firstUniqChar(s string) int {
	byteCount:=make(map[rune]int)
	for _,v:=range s{
		byteCount[v]++
	}
	for i,v:=range s{
		if byteCount[v]==1{
			return i
		}
	}
	return -1
}