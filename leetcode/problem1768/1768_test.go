package problem1768

func mergeAlternately(word1 string, word2 string) string {
    result:=make([]byte,len(word1)+len(word2))
	i,j:=0,0

	for i<len(word1)&&j<len(word2) {
		if i<len(word1){
			result[i+j]=word1[i]
			i++
		}
		if j<len(word2){
			result[i+j]=word2[j]
			j++
		}
	}

	return string(result)
}