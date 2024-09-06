package problem8

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func myAtoi(s string) int {
	negative := false
	skipNonNum := false
	result := float64(0)

	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' || s[i] == '.' {
			break
		}

		if !skipNonNum && !negative && s[i] == '-' {
			negative = true
			skipNonNum = true
		} else if !skipNonNum  && s[i] == '+'{
			skipNonNum = true
		}else if s[i] >= '0' && s[i] <= '9' {
			if negative && result > 0 {
				result = -result
			}
			result *= 10
			if negative {
				result -= float64(s[i] - '0')
			} else {
				result += float64(s[i] - '0')
			}
			result = min(result, math.MaxInt32)
			result = max(result, math.MinInt32)
			skipNonNum = true
		} else if skipNonNum {
			break
		}
	}

	return int(result)
}

func Test8(t *testing.T) {
	// assert.Equal(t, -2147483648, myAtoi("-91283472332"))
	assert.Equal(t, 0, myAtoi("+-12"))
}
