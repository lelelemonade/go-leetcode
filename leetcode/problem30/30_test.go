package problem30

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findSubstring(s string, words []string) []int {
    wordCount:=make(map[string]int)
	wordsLength:=0

	for _,v:=range words{
		wordCount[v]++
		wordsLength+=len(v)
	}

	checkRangeValid:=func(r string)bool{
		currentWordCount:=make(map[string]int)
		for k,v:=range wordCount{
			currentWordCount[k]=v
		}
		wordSize:=len(words[0])
		for i:=0;i<len(r);i+=wordSize {
			if v,e:=currentWordCount[r[i:i+wordSize]];!e||v==0{
				return false
			}else{
				currentWordCount[r[i:i+wordSize]]--
			}
		}

		return true
	}

	result:=[]int{}

	for i:=0;i<=len(s)-wordsLength;i++{
		if checkRangeValid(s[i:i+wordsLength]){
			result = append(result, i)
		}
	}

	return result
}

func Test30(t *testing.T) {
	// assert.Equal(t,[]int{0,9},findSubstring("barfoothefoobarman",[]string{"foo","bar"}))
	assert.Equal(t,[]int{8},findSubstring("wordgoodgoodgoodbestword",[]string{"word","good","best","good"}))
}