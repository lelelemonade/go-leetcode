package problem81

import "fmt"

func search(nums []int, target int) bool {
    var binarySearch func(left int, right int) bool
    binarySearch = func(left int, right int) bool {
        if left > right {
            return false
        }
        
        mid := left + (right - left) / 2

        if nums[mid] == target || nums[left] == target || nums[right] == target {
            return true
        }
        
        if nums[left] == nums[mid] && nums[mid] == nums[right] {
            return binarySearch(left+1, right-1)
        }

        // Left half is sorted
        if nums[left] <= nums[mid] {
            if target >= nums[left] && target < nums[mid] {
                return binarySearch(left, mid-1)
            } else {
                return binarySearch(mid+1, right)
            }
        } else {
            // Right half is sorted
            if target > nums[mid] && target <= nums[right] {
                return binarySearch(mid+1, right)
            } else {
                return binarySearch(left, mid-1)
            }
        }
    }

    return binarySearch(0, len(nums)-1)
}
