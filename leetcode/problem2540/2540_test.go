package problem2540

func getCommon(nums1 []int, nums2 []int) int {
	nums1Iter := 0
	nums2Iter := 0
	for nums1Iter < len(nums1) && nums2Iter < len(nums2) {
		if nums1[nums1Iter]==nums2[nums2Iter]{
			return nums1[nums1Iter]
		}else if nums1[nums1Iter]<nums2[nums2Iter]{
			nums1Iter++
		}else{
			nums2Iter++
		}
	}
	return -1
}
