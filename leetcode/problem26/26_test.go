package problem26

func removeDuplicates(nums []int) int {
    fast,slow:=1,1

	for fast <= len(nums)-1 {
		if nums[fast-1] != nums[fast] {
			nums[slow]=nums[fast]
			slow++
		}
		fast++
	}

	return slow
}