package problem1207

func uniqueOccurrences(arr []int) bool {
    occurrences:=make(map[int]int)
	for _,v:=range arr{
		occurrences[v]++
	}
	occurrencesCount:=make(map[int]struct{})
	for _,v:=range occurrences{
		if _,e:=occurrencesCount[v];e{
			return false
		}else{
			occurrencesCount[v]=struct{}{}
		}
	}
	return true
}