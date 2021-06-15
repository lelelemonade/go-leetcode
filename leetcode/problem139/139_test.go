package problem139

import (
	"fmt"
	"testing"
)

func Test_Problem136(t *testing.T) {
	//s := "catsandog"
	s := "leetcode"
	//testArray := []string{"cats", "dog", "sand", "and", "cat"}
	testArray := []string{"leet", "code", "sand", "and", "cat"}

	fmt.Println(wordBreak(s, testArray))
}
