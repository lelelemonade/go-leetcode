package problem2215

func findDifference(nums1 []int, nums2 []int) [][]int {
    nums1Set:=make(map[int]struct{})
    nums2Set:=make(map[int]struct{})
	for _,v:=range nums1{
		nums1Set[v]=struct{}{}
	}
	for _,v:=range nums2{
		nums2Set[v]=struct{}{}
	}
	result:=make([][]int,0)
	resultNum1:=make([]int,0)
	resultNum2:=make([]int,0)
	for k:=range nums1Set{
		if _,e:=nums2Set[k];!e{
			resultNum1=append(resultNum1, k)
		}
	}
	for k:=range nums2Set{
		if _,e:=nums1Set[k];!e{
			resultNum2=append(resultNum2, k)
		}
	}
	result=append(result, resultNum1)
	result=append(result, resultNum2)
	return result
}