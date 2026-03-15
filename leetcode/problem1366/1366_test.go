package problem1366

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func rankTeams(votes []string) string {
	team := make([]int, 26)

	for _, vote := range votes {
		for i, t := range vote {
			team[t-'A'] += i +1
		}
	}

	scoreToChar := make(map[int]string)
	for t, s := range team {
		if _, ok := scoreToChar[s]; !ok {
			scoreToChar[s] = string(rune('A' + t))
		} else {
			scoreToChar[s] += string(rune('A' + t))
		}

	}

	delete(scoreToChar, 0)

	slices.Sort(team)

	result := ""

	for i := 0; i < len(team); i++ {
		if char, ok := scoreToChar[team[i]]; ok {
			result += char
			delete(scoreToChar, team[i])
		}
	}

	return result
}

func TestRankTeams(t *testing.T) {
	assert.Equal(t, "ACB", rankTeams([]string{"ABC","ACB","ABC","ACB","ACB"}))
}