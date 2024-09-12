package problem1684

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func countConsistentStrings(allowed string, words []string) int {
    allowedSet:=make(map[byte]struct{})

	for _,v:=range allowed{
		allowedSet[byte(v)]=struct{}{}
	}

	result:=0
	wordContinue:
	for _,word:=range words{
		for _,v:=range word{
			if _,e:=allowedSet[byte(v)];!e{
				continue wordContinue
			}
		}
		result++
	}
	return result
}

func Test1684(t *testing.T) {
	assert.Equal(t,2,countConsistentStrings("ab",[]string{"ad","bd","aaab","baa","badab"}))
}