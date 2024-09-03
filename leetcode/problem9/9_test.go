package problem9

func isPalindrome(x int) bool {
	if x<0{
		return false
	}
    reverse:=0
	original:=x

	for original!=0{
		reverse*=10
		reverse+=original%10
		original/=10
	}
	return reverse==x
}