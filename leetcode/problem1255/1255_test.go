package problem1255

import (
	"github.com/stretchr/testify/assert"
	"maps"
	"testing"
)

func Test1255(t *testing.T) {
	assert.Equal(t, 27, maxScoreWords([]string{"xxxz", "ax", "bx", "cx"}, []byte{'z', 'a', 'b', 'c', 'x', 'x', 'x'}, []int{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10}))
	assert.Equal(t, 23,
		maxScoreWords(
			[]string{"dog", "cat", "dad", "good"},
			[]byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'},
			[]int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		),
	)
}

func maxScoreWords(words []string, letters []byte, score []int) int {
	totalChoiceCount := 1 << len(words)
	letterCountMap := make(map[byte]int)
	for _, v := range letters {
		letterCountMap[v] += 1
	}

	letterScoreMap := make(map[byte]int)
	for i, v := range score {
		letterScoreMap[byte('a'+i)] = v
	}

	currentLetterCount := make(map[byte]int)
	maxScore := 0
	currentScore := 0
	for i := 0; i < totalChoiceCount; i++ {
		for j := 0; j < len(words); j++ {
			if i>>j == 0 {
				break
			}
			if (i>>j)&1 == 1 {
				for _, v := range words[j] {
					currentLetterCount[byte(v)] += 1
					currentScore += letterScoreMap[byte(v)]
				}
			}
		}

		valid := true

		for k, v := range currentLetterCount {
			if v > letterCountMap[k] {
				valid = false
				break
			}
		}
		if valid {
			maxScore = max(maxScore, currentScore)
		}
		currentScore = 0
		maps.DeleteFunc(currentLetterCount, func(b byte, i int) bool {
			return true
		})
	}

	return maxScore
}
