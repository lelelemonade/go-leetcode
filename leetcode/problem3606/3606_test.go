package problem3606

import (
	"slices"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	resultGroup := make([][]string, 4)

	for i, c := range code {
		isCode := func(s string) bool {
			for _, j := range s {
				if j != '_' && !unicode.IsLetter(j) && !unicode.IsDigit(j) {
					return false
				}
			}
			return true
		}
		if !isActive[i] || c == "" || !isCode(c) {
			continue
		}

		switch businessLine[i] {
		case "electronics":
			resultGroup[0] = append(resultGroup[0], c)
		case "grocery":
			resultGroup[1] = append(resultGroup[1], c)
		case "pharmacy":
			resultGroup[2] = append(resultGroup[2], c)
		case "restaurant":
			resultGroup[3] = append(resultGroup[3], c)
		}
	}

	var result []string

	for _, r := range resultGroup {
		slices.Sort(r)
		result = append(result, r...)
	}

	return result
}

func TestValidateCoupons(t *testing.T) {
	code := []string{"SAVE20", "", "PHARMA5", "SAVE@20"}
	businessLine := []string{"restaurant", "grocery", "pharmacy", "restaurant"}
	isActive := []bool{true, true, true, true}

	expect := []string{"PHARMA5", "SAVE20"}

	assert.Equal(t, expect, validateCoupons(code, businessLine, isActive))
}
