package problem3083

func isSubstringPresent(s string) bool {
	subMap:=make(map[string]struct{})

	for i:=0;i<len(s)-1;i++{
		subMap[s[i:i+2]]=struct{}{}
	}

	sReverse:=""
	for i:=len(s)-1;i>=0;i--{
		sReverse+=string(s[i])
	}

	for i:=0;i<len(sReverse)-1;i++{
		if _,e:=subMap[sReverse[i:i+2]];e{
			return true
		}
	}
	return false
}