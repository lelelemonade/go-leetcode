package problem258

func addDigits(num int) int {
    for num/10!=0{
		current:=0
		for num!=0{
			current+=num%10
			num/=10
		}
		num=current
	}
	return num
}