package problem1909

func canBeIncreasing(nums []int) bool {
    n := len(nums)
    count := 0 

    for i := 1; i < n; i++ {
        if nums[i] <= nums[i-1] {
            count++
            if count > 1 {
                return false
            }
            if i > 1 && nums[i] <= nums[i-2] {
                nums[i] = nums[i-1]
            }
        }
    }

    return true
}