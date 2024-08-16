package problem202

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isHappy(n int) bool {
	iterBefore := make(map[int]struct{}, 0)

	for {
		nextN := 0
		for n != 0 {
			nextN += (n % 10) * (n % 10)
			n = n / 10
		}
		if nextN == 1 {
			return true
		}
		if _, e := iterBefore[nextN]; e {
			return false
		}
		iterBefore[nextN] = struct{}{}
		n = nextN
	}
}

func Test19(t *testing.T) {
	assert.Equal(t, false, isHappy(2))
}
