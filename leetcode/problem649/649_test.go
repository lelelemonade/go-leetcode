package problem649

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func predictPartyVictory(senate string) string {
	radiantOutNum:=0
	direOutNum:=0
	senateArray := []rune(senate)
	senateArray = append(senateArray, ' ')
	outCount := 0

	for len(senateArray) > 0 && outCount < len(senate)+1 {
		head := senateArray[0]
		if head == ' ' {
			outCount = 0
		}else{
			outCount++
		}
		senateArray = senateArray[1:]
		if head == 'R' && radiantOutNum > 0 {
			radiantOutNum--
			senateArray = append(senateArray, ' ')
		} else if head == 'D' && direOutNum>0 {
			direOutNum--
			senateArray = append(senateArray, ' ')
		} else if head == 'R' {
			direOutNum++
			senateArray = append(senateArray, head)
		}else if head == 'D' {
			radiantOutNum++
			senateArray = append(senateArray, head)
		}
	}

	if senateArray[0] == 'D' {
		return "Dire"
	}

	return "Radiant"
}

func Test649(t *testing.T) {
	assert.Equal(t, "Radiant", predictPartyVictory("RD"))
	assert.Equal(t, "Dire", predictPartyVictory("RDD"))
	assert.Equal(t, "Dire", predictPartyVictory("DDRRR"))
}
