package problem556

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func Test556(t *testing.T) {
	//fmt.Println(nextGreaterElement(21))
	fmt.Println(nextGreaterElement(1234))
	//fmt.Println(nextGreaterElement(12))
}

func nextGreaterElement(n int) int {
	arrayN := []byte(strconv.Itoa(n))

	if len(arrayN) < 2 {
		return -1
	}

	left := len(arrayN) - 2

	for ; left >= 0; left-- {
		if arrayN[left] < arrayN[left+1] {
			break
		}
	}

	if left == -1 {
		return -1
	}

	right := len(arrayN) - 1
	for ; right > left; right-- {
		if arrayN[right] > arrayN[left] {
			break
		}
	}

	arrayN[left], arrayN[right] = arrayN[right], arrayN[left]

	i := left + 1
	j := len(arrayN) - 1
	for i != j && i < j {
		arrayN[i], arrayN[j] = arrayN[j], arrayN[i]
		i++
		j--
	}

	result, err := strconv.Atoi(string(arrayN))

	if err != nil || result == n || result > math.MaxInt32 {
		return -1
	}

	return result

	//searching := true
	//left := len(arrayN) - 1
	//
	//for searching && left > 0 {
	//	for right := left - 1; right >= 0; right-- {
	//		if arrayN[right] < arrayN[left] {
	//			temp := arrayN[left]
	//			iter := left - 1
	//			for ; iter >= right; iter-- {
	//				arrayN[iter+1] = arrayN[iter]
	//			}
	//			arrayN[right] = temp
	//			searching = false
	//			break
	//		}
	//	}
	//	left--
	//}
	//
	//result, err := strconv.Atoi(string(arrayN))
	//if err != nil || result == n {
	//	return -1
	//} else {
	//	return result
	//}
}
