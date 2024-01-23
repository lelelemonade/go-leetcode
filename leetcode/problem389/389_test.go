package problem389

import (
	"fmt"
	"testing"
)

func Test_389(t *testing.T) {
	fmt.Println(findTheDifference("afdfsdfdsfs", "aafdfsdfdsfs"))
}

func findTheDifference(s string, t string) byte {
	total := 1
	for i := 0; i < len(s); i++ {
		total = total * int(t[i]) / int(s[i])
	}

	total = total * int(t[len(t)-1])

	//for i := 0; i < len(s); i++ {
	//	total = total / int(s[i])
	//}

	return byte(total)
}
