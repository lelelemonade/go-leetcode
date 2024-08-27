package problem1456

func maxVowels(s string, k int) int {
	isVowels:=func(b byte)bool{
		return b=='a'||b=='e'||b=='i'||b=='o'||b=='u'
	}

    windowVowelCount:=0
	for i:=0;i<k;i++{
		if isVowels(s[i]){
			windowVowelCount++
		}
	}
	result:=windowVowelCount

	for i:=k;i<len(s);i++{
		if isVowels(s[i]){
			windowVowelCount++
		}
		if isVowels(s[i-k]){
			windowVowelCount--
		}
		result=max(result,windowVowelCount)
	}

	return result
}