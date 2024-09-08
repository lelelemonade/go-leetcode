package problem29

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	sign := 1
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		sign = -1
	}

	dividendAbs := abs(dividend)
	divisorAbs := abs(divisor)

	quotient := 0

	for dividendAbs >= divisorAbs {
		tempDivisor, multiple := divisorAbs, 1
		for dividendAbs >= (tempDivisor << 1) {
			tempDivisor <<= 1
			multiple <<= 1
		}
		dividendAbs -= tempDivisor
		quotient += multiple
	}

	if sign == -1 {
		return -quotient
	}
	return quotient
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Test29(t *testing.T) {
	assert.Equal(t,3,divide(10,3))
}