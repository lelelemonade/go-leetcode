package problem735

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	checkCollide := func(asteroid int) (isAsteroidWin bool) {
		if len(stack) == 0 || stack[len(stack)-1]<0||asteroid>0 {
			stack = append(stack, asteroid)
			return false
		}
		if abs(stack[len(stack)-1]) == abs(asteroid) {
			stack = stack[:len(stack)-1]
			return false
		} else if abs(stack[len(stack)-1]) < abs(asteroid) {
			stack = stack[:len(stack)-1]
			return true
		} else {
			return false
		}
	}

	for _, v := range asteroids {
		for checkCollide(v) {

		}
	}

	return stack
}

func Test735(t *testing.T){
	assert.Equal(t,[]int{5,10},asteroidCollision([]int{5,10,-5}))
	assert.Equal(t,[]int{-2,-1,1,2},asteroidCollision([]int{-2,-1,1,2}))
}
