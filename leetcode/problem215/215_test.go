package problem215

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test215(t *testing.T) {
	//assert.Equal(t, 4, findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	//assert.Equal(t, 2, findKthLargest([]int{2, 1}, 1))
	//assert.Equal(t, 1, findKthLargest([]int{1}, 1))
	//assert.Equal(t, 5, findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	assert.Equal(t, 6, findKthLargest([]int{7, 6, 5, 4, 3, 2, 1}, 2))
}

func findKthLargest(nums []int, k int) int {
    var bucket [20001]int

    for _,v := range nums {
        bucket[v+10000]++
    }

    count:=0
    for i:=20000;i>=0;i--{
        count+= bucket[i]
        if count >= k{
            return i-10000
        }
    }

    return -1
}
