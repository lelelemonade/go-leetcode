package problem2614

func diagonalPrime(nums [][]int) int {
	isPrime:=func(a int)bool{
		for i:=2;i*i<=a;i++{
			if a%i==0{
				return false
			}
		}
		return a >= 2
	}
	result:=0

	for i:=0;i<len(nums);i++{
		if nums[i][i]>result&& isPrime(nums[i][i]){
			result=nums[i][i]
		}
		if nums[i][len(nums)-1-i]>result&& isPrime(nums[i][len(nums)-1-i]){
			result=nums[i][len(nums)-1-i]
		}
	}

	return result
}