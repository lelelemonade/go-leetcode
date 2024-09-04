package problem172

func trailingZeroes(n int) int {
	var result int
	if n >= 3125 {
		result += n / 3125
	}
	if n >= 625 {
		result += n / 625
	}
	if n >= 125 {
		result += n / 125
	}
	if n >= 25 {
		result += n / 25
	}
	result += n / 5
	return result
}
