package problem88

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func merge(nums1 []int, m int, nums2 []int, n int)  {
    iter1,iter2:=m-1,n-1
	idx:=m+n-1

	for iter1>=0&&iter2>=0{
		if nums1[iter1]<nums2[iter2]{
			nums1[idx]=nums2[iter2]
			iter2--
		}else{
			nums1[idx]=nums1[iter1]
			iter1--
		}
		idx--
	}
	for iter2 >=0 {
		nums1[idx]=nums2[iter2]
		iter2--
		idx--
	}
}

func Test88(t *testing.T) {
	testCase1Num1:=[]int{1,2,3,0,0,0}
	testCase1M:=3
	testCase1Num2:=[]int{2,5,6}
	testCase1N:=3

	merge(testCase1Num1,testCase1M,testCase1Num2,testCase1N)

	assert.Equal(t,[]int{1,2,2,3,5,6},testCase1Num1)
}