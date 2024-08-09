package problem860

import (
	"fmt"
	"testing"
)

func Test860(t *testing.T) {
	fmt.Println(lemonadeChange([]int{5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5}))
}

func lemonadeChange(bills []int) bool {
	fiveDollarCount := 0
	tenDollarCount := 0
	twentyDollarCount := 0

	for i := 0; i < len(bills); i++ {
		if bills[i] == 20 {
			if tenDollarCount >= 1 {
				tenDollarCount--
				fiveDollarCount--
			} else {
				fiveDollarCount -= 3
			}
			twentyDollarCount++
			if fiveDollarCount < 0 {
				return false
			}
		} else if bills[i] == 10 {
			fiveDollarCount--
			tenDollarCount++
			if fiveDollarCount < 0 {
				return false
			}
		} else if bills[i] == 5 {
			fiveDollarCount++
		}
	}

	return true
}
