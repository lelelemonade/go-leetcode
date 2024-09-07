package problem507

func checkPerfectNumber(num int) bool {
	if num ==1{
		return false
	}
    result:=1

	for i:=2;i*i<num;i++{
		if num%i==0{
			result+=num/i+i
		}
	}

	return result==num
}