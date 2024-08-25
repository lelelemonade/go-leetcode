package problem2147

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func numberOfWays(corridor string) int {
	seatCount := 0
	for _, v := range corridor {
		if v == 'S' {
			seatCount++
		}
	}
	if seatCount == 0 || seatCount%2 == 1 {
		return 0
	}
	result:=1
	seatCount=0
    for i,v := range corridor {
		if v=='S'{
			if seatCount==2{
				plantCount:=i-1
				for corridor[plantCount]!='S'{
					plantCount--
				}
				result=(result*(i-plantCount))%1000000007
				fmt.Println(result)
				seatCount=0
			}
			seatCount+=1
		}
	}
	return result
}

func Test2147(t *testing.T) {
	// assert.Equal(t, 3, numberOfWays("SSPPSPS"))
	// assert.Equal(t, 1, numberOfWays("PPSPSP"))
	// assert.Equal(t, 0, numberOfWays("S"))
	// assert.Equal(t, 1, numberOfWays("SPPSSSSPPS"))
	assert.Equal(t, 919999993, numberOfWays("PPPPPSPPSPPSPPPSPPPPSPPPPSPPPPSPPSPPPSPSPPPSPSPPPSPSPPPSPSPPPPSPPPPSPPPSPPSPPPPSPSPPPPSPSPPPPSPSPPPSPPSPPPPSPSPSS"))
}
