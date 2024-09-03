package problem191

func hammingWeight(n int) int {
    result:=0

	for n!=0{
		result+=n&1
		n>>=1
	}
	return result
}