package problem383

func canConstruct(ransomNote string, magazine string) bool {
    magazineCount := make(map[rune]int)

	for _,v := range magazine {
		if count,exist:=magazineCount[v];exist{
			magazineCount[v]=count+1
		}else{
			magazineCount[v]=1
		}
	}

	ransomNoteCount := make(map[rune]int)
	for _, v := range ransomNote {
		if count,exist:=ransomNoteCount[v];exist{
			ransomNoteCount[v]=count+1
		}else{
			ransomNoteCount[v]=1
		}
	}

	for k,v := range ransomNoteCount {
		if magazineCount[k]<v{
			return false
		}
	}

	return true
}

func canConstruct2(ransomNote, magazine string) bool {
    if len(ransomNote) > len(magazine) {
        return false
    }
    cnt := [26]int{}
    for _, ch := range magazine {
        cnt[ch-'a']++
    }
    for _, ch := range ransomNote {
        cnt[ch-'a']--
        if cnt[ch-'a'] < 0 {
            return false
        }
    }
    return true
}