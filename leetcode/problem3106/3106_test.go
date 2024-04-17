package problem3106

import (
	"fmt"
	"testing"
)

func Test3106(t *testing.T) {
	// fmt.Println(getSmallestString("testmeplease", 3))
	// fmt.Println(getSmallestString("zbbz", 3))
	fmt.Println(getSmallestString("xaxcd", 4))
	fmt.Println(getSmallestString("lol", 0))
}

func getSmallestString(s string, k int) string {
	byteArray := []byte(s)

	for i := 0; i < len(byteArray) && k > 0; i++ {
		if distance := int(min(byteArray[i]-'a', 'z'+1-byteArray[i])); distance <= k {
			k -= distance
			byteArray[i] = 'a'
		} else {
			if byteArray[i]-'a' >= byte(k) {
				byteArray[i] -= byte(k)
				break
			} else {
				k -= int(byteArray[i] - 'a')
				byteArray[i] = 'a'

			}
		}
	}

	return string(byteArray)
}
