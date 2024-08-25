package problem2165

import (
	"slices"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func smallestNumber(num int64) int64 {
	var numString string
	if num <0{
		numString = strconv.FormatInt(-num,10)
	}else{
		numString = strconv.FormatInt(num,10)
	}
	numSlice:=make([]rune,0)
	for _,v:=range numString{
		numSlice=append(numSlice, v)
	}

	if num <0{
		slices.SortFunc(numSlice,func(a rune,b rune) int {return int(b-a)})

		result,err:=strconv.ParseInt(string(numSlice), 10, 64)

		if err != nil {
			panic(err)
		}
	
		return -result
	}else{
		slices.Sort(numSlice)
	}

	for i,v:=range numSlice{
		if v!='0'{
			numSlice[i],numSlice[0]=numSlice[0],numSlice[i]
			break
		}
	}

	result,err:=strconv.ParseInt(string(numSlice), 10, 64)

	if err != nil {
		panic(err)
	}

	return result
}

func Test2165(t *testing.T) {
	assert.Equal(t,int64(103),smallestNumber(301))
	assert.Equal(t,int64(-7650),smallestNumber(-7605))
}