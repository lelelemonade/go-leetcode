package problem1813

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	words1 := strings.Split(sentence1, " ")
	words2 := strings.Split(sentence2, " ")
	left, right := 0, 0
	n := len(words1)
	m := len(words2)
	for left < min(n, m) && words1[left] == words2[left] {
		left++
	}

	for right < min(n-left, m-left) && words1[n-1-right] == words2[m-1-right] {
		right++
	}
	return left+right == min(n, m)
}

func Test1813(t *testing.T) {
	assert.Equal(t,true,areSentencesSimilar("MTNhOb epj zqI dR","MTNhOb epj dR"))
}