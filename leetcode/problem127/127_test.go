package problem127

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test127(t *testing.T) {
	beginWord1 := "hit"
	endWord1 := "cog"
	wordList1 := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	assert.Equal(t, 5, ladderLength(beginWord1, endWord1, wordList1))

}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	stepMap := make(map[string]*Step)

	stepMap[beginWord] = &Step{
		word:         beginWord,
		adjacentStep: make([]*Step, 0),
	}

	for _, word := range wordList {
		stepMap[word] = &Step{
			word:         word,
			adjacentStep: make([]*Step, 0),
		}
	}

	for k, v := range stepMap {
		for _, word := range wordList {
			if adjacent(k, word) {
				v.adjacentStep = append(v.adjacentStep, stepMap[word])
			}
		}
	}

	//bfs
	stack := make([]*Step, 0)
	stepMap[beginWord].depth = 1
	stack = append(stack, stepMap[beginWord])
	for len(stack) != 0 {
		if stack[0].word == endWord {
			return stack[0].depth
		}

		for _, v := range stack[0].adjacentStep {
			if v.depth == 0 {
				v.depth = stack[0].depth + 1
				stack = append(stack, v)
			}
		}

		stack = stack[1:]
	}

	return 0
}

func adjacent(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	diffCount := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diffCount++
		}
	}
	return diffCount == 1
}

type Step struct {
	word         string
	depth        int
	adjacentStep []*Step
}
