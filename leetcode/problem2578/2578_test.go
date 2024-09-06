package problem2578

import "strconv"

func splitNum(num int) int {
	var numCount [9]int
	for _, v := range strconv.Itoa(num) {
		if v == '0' {
			continue
		}
		numCount[v-'1']++
	}

	num1 := 0
	num2 := 0
	addToNum1 := true

	for i := 0; i < 9; i++ {
		for numCount[i] != 0 {
			if addToNum1 {
				num1 *= 10
				num1 += i + 1
				addToNum1=false
			} else {
				num2 *= 10
				num2 += i + 1
				addToNum1=true
			}
			numCount[i]--
		}
	}

	return num1+num2
}
