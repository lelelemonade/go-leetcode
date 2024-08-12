package problem12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func intToRoman(num int) string {
	romanSymbolMap := make(map[int]string)
	romanSymbolMap[1] = "I"
	romanSymbolMap[2] = "II"
	romanSymbolMap[3] = "III"
	romanSymbolMap[4] = "IV"
	romanSymbolMap[5] = "V"
	romanSymbolMap[6] = "VI"
	romanSymbolMap[7] = "VII"
	romanSymbolMap[8] = "VIII"
	romanSymbolMap[9] = "IX"
	romanSymbolMap[10] = "X"
	romanSymbolMap[20] = "XX"
	romanSymbolMap[30] = "XXX"
	romanSymbolMap[40] = "XL"
	romanSymbolMap[50] = "L"
	romanSymbolMap[60] = "LX"
	romanSymbolMap[70] = "LXX"
	romanSymbolMap[80] = "LXXX"
	romanSymbolMap[90] = "XC"
	romanSymbolMap[100] = "C"
	romanSymbolMap[200] = "CC"
	romanSymbolMap[300] = "CCC"
	romanSymbolMap[400] = "CD"
	romanSymbolMap[500] = "D"
	romanSymbolMap[600] = "DC"
	romanSymbolMap[700] = "DCC"
	romanSymbolMap[800] = "DCCC"
	romanSymbolMap[900] = "CM"
	romanSymbolMap[1000] = "M"
	romanSymbolMap[2000] = "MM"
	romanSymbolMap[3000] = "MMM"

	return romanSymbolMap[num-num%1000]+romanSymbolMap[num%1000-num%100]+romanSymbolMap[num%100-num%10]+romanSymbolMap[num%10]
}

func Test12(t *testing.T) {
	assert.Equal(t,"MMMDCCXLIX",intToRoman(3749))
}