package problem2929

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func distributeCandies(n int, limit int) int64 {
	result:=int64(0)

	for i:=0;i<=min(limit, n);i++{
		if n-i>2*limit{
			continue
		}
		if n-i<=limit{
			result+=int64(n-i+1)
		}else{
			result+=int64(2*limit-(n-i)+1)
		}
	}

	return result
}

func Test2929(t *testing.T) {
	assert.Equal(t,3,distributeCandies(5,2))
}