package problem13

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func romanToInt(s string) int {
	result := 0

	idx := 0

	for idx < len(s) {
		if len(s)-idx > 1 {
			switch s[idx : idx+2] {
			case "IV":
				result += 4
				idx += 2
				continue
			case "IX":
				result += 9
				idx += 2
				continue
			case "XL":
				result += 40
				idx += 2
				continue
			case "XC":
				result += 90
				idx += 2
				continue
			case "CD":
				result += 400
				idx += 2
				continue
			case "CM":
				result += 900
				idx += 2
				continue
			}
		}

		switch s[idx : idx+1] {
		case "I":
			result += 1
		case "V":
			result += 5
		case "X":
			result += 10
		case "L":
			result += 50
		case "C":
			result += 100
		case "D":
			result += 500
		case "M":
			result += 1000
		}
		idx += 1
	}

	return result
}

func Test13(t *testing.T) {
	assert.Equal(t,58,romanToInt("LVIII"))
}