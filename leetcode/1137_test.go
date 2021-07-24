package leetcode

import (
	"fmt"
	"testing"
)

func tribonacci(n int) int {
	a, b, c := 0, 1, 1
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	ans := 0
	for i := 3; i <= n; i++ {
		ans = a + b + c
		a = b
		b = c
		c = ans
	}
	return ans

}

func Test1137(t *testing.T) {
	fmt.Println(tribonacci(25))
}
