package problem345

func reverseVowels(s string) string {
    vowels:=make([]byte,0)
	sArray:=[]byte(s)

	isVowel:=func(b byte)bool{
		return b=='a'||
		b=='e'||
		b=='i'||
		b=='o'||
		b=='u'||
		b=='A'||
		b=='E'||
		b=='I'||
		b=='O'||
		b=='U'
	}

	for _,v:=range sArray{
		if isVowel(v){
			vowels = append(vowels, v)
		}
	}
	vowelsIter:=len(vowels)-1
	for i:=0;i<len(s);i++{
		if isVowel(sArray[i]){
			sArray[i]=byte(vowels[vowelsIter])
			vowelsIter--
		}
	}

	return string(sArray)
}

