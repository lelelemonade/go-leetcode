package problem169

func majorityElement(nums []int) int {
	countMap := make(map[int]int)
	majority := 0

	for _, v := range nums {
		if _, exist := countMap[v]; exist {
			countMap[v] = countMap[v] + 1
		} else {
			countMap[v] = 1
		}

		if countMap[majority]<countMap[v]{
			majority=v
		}
	}
	return majority
}
