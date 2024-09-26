package problem2416

func sumPrefixScores(words []string) []int {
    prefixSum:=make(map[string]int)

	for _,word:=range words{
		for i := range word {
			prefixSum[word[:i+1]]++
		}
	}

	results:=make([]int,len(words))

	for i,word:=range words{
		result:=0

		for i := range word {
			result+=prefixSum[word[:i+1]]
		}

		results[i]=result
	}

	return results
}