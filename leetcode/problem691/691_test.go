package problem691

import (
	"github.com/stretchr/testify/assert"
	"maps"
	"math"
	"math/bits"
	"testing"
)

func Test691(t *testing.T) {
	assert.Equal(t, 3, minStickers([]string{"with", "example", "science"}, "thehat"))
}

func minStickers(stickers []string, target string) int {
	stickerLettersMap := make([]map[byte]int, 0)
	targetLetterMap := make(map[byte]int)

	for i := 0; i < len(target); i++ {
		targetLetterMap[target[i]] += 1
	}

	for _, sticker := range stickers {
		stickerLetterMap := make(map[byte]int)

		shouldSkip := true

		for i := 0; i < len(sticker); i++ {
			stickerLetterMap[sticker[i]] += 1
			if targetLetterMap[sticker[i]] != 0 {
				shouldSkip = false
			}
		}
		if !shouldSkip {
			stickerLettersMap = append(stickerLettersMap, stickerLetterMap)
		}
	}

	minNum := math.MaxInt
	currentLetterMap := make(map[byte]int)
	currentNumberMap := make(map[int]int)
	for i := 0; i < 1<<len(stickerLettersMap); i++ {
		for j := 0; i>>j > 0; j++ {
			if i>>j&1 == 1 {
				for k, v := range stickerLettersMap[j] {
					currentLetterMap[k] += v
				}
			}
		}

		valid := true
		for k, v := range targetLetterMap {
			if currentLetterMap[k] == 0 {
				valid = false
				break
			} else {
				currentNumberMap[] = v / currentLetterMap[k]
			}
		}
		if valid {
			minNum = min(minNum, bits.OnesCount(uint(i)))
		}

		maps.DeleteFunc(currentLetterMap, func(b byte, i int) bool {
			return true
		})
		maps.DeleteFunc(currentNumberMap, func(b int, i int) bool {
			return true
		})
	}

	if minNum == math.MaxInt {
		return -1
	}
	return minNum
}
