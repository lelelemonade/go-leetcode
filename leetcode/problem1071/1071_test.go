package problem1071

func gcdOfStrings(str1 string, str2 string) string {
	minLength:=min(len(str1),len(str2))
	for i:=0;i<minLength;i++{
		if str1[i]!=str2[i]{
			return ""
		}
	}

    if len(str1)==len(str2){
		return str1
	}else if len(str1)<len(str2){
		return gcdOfStrings(str1,str2[minLength:])
	}else{
		return gcdOfStrings(str2,str1[minLength:])
	}
}