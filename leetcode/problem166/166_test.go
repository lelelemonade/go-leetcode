package problem166

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func fractionToDecimal(numerator int, denominator int) string {
	result := ""
	if numerator == 0 {
		return "0"
	}
	if numerator*denominator < 0 {
		result += "-"
		if numerator<0{
			numerator=-numerator
		}
		if denominator<0{
			denominator=-denominator
		}
	}
	result += strconv.Itoa(numerator / denominator)
	if numerator % denominator==0{
		return result
	}

	numerator %= denominator
	result += "."
	indexMap := map[int]int{}

	for numerator != 0 && indexMap[numerator] == 0 {
		indexMap[numerator] = len(result)
		result += strconv.Itoa(numerator * 10 / denominator)
		numerator = numerator * 10 % denominator
	}

	if numerator > 0 {
		result = result[:indexMap[numerator]] + "(" + result[indexMap[numerator]:] + ")"
	}

	return result
}

func Test166(t *testing.T) {
	assert.Equal(t,"0.1(6)",fractionToDecimal(1,6))
	assert.Equal(t,"0.5",fractionToDecimal(1,2))
	assert.Equal(t,"2",fractionToDecimal(2,1))
	assert.Equal(t, "0.(012)", fractionToDecimal(4, 333))
}
