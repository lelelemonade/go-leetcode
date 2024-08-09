package problem55

func canJump(nums []int) bool {
    maxJumpIdx := 0
    for i := 0; i <len(nums);i++ {
        if i > maxJumpIdx {
            return false
        }

        maxJumpIdx = max(nums[i]+i,maxJumpIdx)
    }
    return true
}