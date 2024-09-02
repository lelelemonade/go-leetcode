package problem3016

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func minimumPushes(word string) int {
    wordCount:=make([]int,26)
	for _,v:=range word{
		wordCount[v-'a']++
	}

	wordIdx:=make([]int,26)
	for i,_:=range wordIdx{
		wordIdx[i]=i
	}

	slices.SortFunc(wordIdx, func(i,j int)int{
		return wordCount[j]-wordCount[i]
	})
	
	result :=0
	wordSize:=0

	for i:=0;i<len(wordIdx);i++{
		if wordCount[wordIdx[i]]==0{
			continue
		}
		result+=wordCount[wordIdx[i]]*((wordSize/8)+1)
		wordSize++
	}

	return result
}

func Test3016(t *testing.T) {
	assert.Equal(t,32,minimumPushes("abzaqsqcyrbzsrvamylmyxdjl"))
}