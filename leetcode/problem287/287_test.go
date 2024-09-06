package problem287

func findDuplicate(nums []int) int {
	fast, slow := 0, 0
	for slow,fast = nums[slow],nums[nums[fast]];fast != slow; {
		slow,fast = nums[slow],nums[nums[fast]]
	}
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
