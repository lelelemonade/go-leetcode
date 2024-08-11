package problem42

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func trap(height []int) int {
    leftMax := make([]int, len(height))

	leftOfMax := 0
	for i := 0; i < len(height); i++ {
		leftMax[i]=leftOfMax
		leftOfMax=max(leftOfMax,height[i])
	}

	rightMax := make([]int, len(height))
	rightOfMax :=0
	for i :=len(height)-1;i>=0;i-- {
		rightMax[i]=rightOfMax
		rightOfMax=max(rightOfMax,height[i])
	}

	trapWater := 0
	for i:=1;i<len(height)-1;i++ {
		trapWater+=max(min(leftMax[i],rightMax[i])-height[i],0)
	}

	return trapWater
}

func Test42(t *testing.T) {
	assert.Equal(t,6,trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
}