package problem47

func permuteUnique(nums []int) [][]int {
	numsSet := make(map[int]struct{})

	results := make([][]int, 0)

	backtrack(numsSet, []int{}, &results)

	return results

}

func backtrack(numsSet map[int]struct{}, chosen []int, results *[][]int) {

}
