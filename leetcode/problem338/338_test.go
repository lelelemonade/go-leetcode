package problem338

import "math/bits"

func countBits(n int) []int {
    result:=make([]int,n+1)

	for i:=0;i<=n;i++{
		result[i]=bits.OnesCount(uint(i))
	}

	return result
}