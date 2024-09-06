package problem168

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func convertToTitle(columnNumber int) string {
	result:=[]byte{}
    for columnNumber!=0{
		mod:=(columnNumber+25)%26
		result=append(result, byte(mod+'A'))
		columnNumber=(columnNumber-mod)/26
	}
	slices.Reverse(result)
	return string(result)
}

func Test168(t *testing.T){
	assert.Equal(t,"ZY",convertToTitle(701))
}

