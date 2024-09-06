package problem350

func intersect(nums1 []int, nums2 []int) []int {
    numCount:=make(map[int]int)

	for _,v:=range nums1{
		numCount[v]++
	}
	result:=make([]int,0)
	for _,v:=range nums2{
		if numCount[v]>0{
			result = append(result, v)
			numCount[v]--
		}
	}
	return result
}