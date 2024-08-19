package problem264

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func nthUglyNumberNew(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(x2, x3, x5)
		if dp[i] == x2 {
			p2++
		}else if dp[i] == x3 {
			p3++
		}else if dp[i] == x5 {
			p5++
		}
	}
	return dp[n]
}

func nthUglyNumber(n int) int {
	uglyMap := make(map[int]struct{})
	uglyIter := 1

	for len(uglyMap) < n {
		i := uglyIter

		for i != 1 && i != 2 && i != 3 && i != 5 {
			if _,e:=uglyMap[i];e{
				uglyMap[uglyIter] = struct{}{}
				break
			}

			if i%2 == 0 {
				i = i / 2
			} else if i%3 == 0 {
				i = i / 3
			} else if i%5 == 0 {
				i = i / 5
			} else {
				break
			}
		}

		if i == 1 || i == 2 || i == 3 || i == 5 {
			uglyMap[uglyIter] = struct{}{}
		}
		uglyIter++
	}

	return uglyIter - 1
}

func Test264(t *testing.T) {
	assert.Equal(t, 12, nthUglyNumber(10))
	assert.Equal(t, 402653184, nthUglyNumber(1352))
}
