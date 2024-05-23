package problem494

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test494(t *testing.T) {
	assert.Equal(t, 5, findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	assert.Equal(t, 5, findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	assert.Equal(t, 256, findTargetSumWays([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1))
}

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if int(math.Abs(float64(target))) > sum {
		return 0
	}
	if (sum+target)%2 == 1 {
		return 0
	}
	// 计算背包大小
	bag := (sum + target) / 2
	// 定义dp数组
	dp := make([]int, bag+1)
	// 初始化
	dp[0] = 1
	// 遍历顺序
	for i := 0; i < len(nums); i++ {
		for j := bag; j >= nums[i]; j-- {
			//推导公式
			dp[j] += dp[j-nums[i]]
			//fmt.Println(dp)
		}
	}
	return dp[bag]
}
