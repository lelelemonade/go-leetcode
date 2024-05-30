package problem1742

func countBalls(lowLimit int, highLimit int) int {
	boxMap := make(map[int]int)
	result := 0

	for i := lowLimit; i <= highLimit; i++ {
		digitSum := countDigitSum(i)
		boxMap[digitSum]++
		result = max(boxMap[digitSum], result)
	}
	return result
}

func countDigitSum(num int) (result int) {
	for num != 0 {
		result += num % 10
		num = num / 10
	}
	return
}
