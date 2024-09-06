package problem2566

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func minMaxDifference(num int) int {
    biggestString:=strconv.Itoa(num)
    smallestString:=strconv.Itoa(num)
	iter:=0
	for iter<len(biggestString)&& biggestString[iter]=='9'{
		iter++
	}
	biggest:=num
	if iter==len(biggestString){
		biggest=num
	}else{
		biggest,_=strconv.Atoi(strings.ReplaceAll(biggestString,string(biggestString[iter]),"9"))
	}

	smallest,e:=strconv.Atoi(strings.ReplaceAll(smallestString,string(smallestString[0]),"0"))
	if e!=nil{
		panic(e)
	}

	return biggest-smallest
}

func Test2566(t *testing.T) {
	assert.Equal(t,99009,minMaxDifference(11891))
}