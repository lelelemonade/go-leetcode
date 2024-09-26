package problem904

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func totalFruit(fruits []int) int {
    fruitNum:=make(map[int]int)
	left:=0
	maxFruits:=0
	for right:=0;right<len(fruits);right++{
		fruitNum[fruits[right]]++
		for len(fruitNum)>2{
			fruitNum[fruits[left]]--
			if fruitNum[fruits[left]]==0{
				delete(fruitNum,fruits[left])
			}
			left++
		}
		maxFruits=max(maxFruits,right-left+1)
	}
	return maxFruits
}

func Test904(t *testing.T) {
	assert.Equal(t,3,totalFruit([]int{1,2,1}))
}