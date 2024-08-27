package problem1657

func closeStrings(word1 string, word2 string) bool {
	if len(word1)!=len(word2){
		return false
	}
    word1Count:=make(map[rune]int)
    word2Count:=make(map[rune]int)
	for _,v:=range word1{
		word1Count[v]++
	}
	for _,v:=range word2{
		word2Count[v]++
	}
	if len(word1Count)!=len(word2Count){
		return false
	}
	for k:=range word1Count{
		if _,e:=word2Count[k];!e{
			return false
		}
	}
	wordCountCount:=make(map[int]int)
	for _,v:=range word1Count{
		wordCountCount[v]++
	}
	for _,v:=range word2Count{
		wordCountCount[v]--
	}
	for _,v:=range wordCountCount{
		if v!=0{
			return false
		}
	}
	return true
}