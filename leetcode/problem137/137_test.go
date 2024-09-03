package problem137

func singleNumber(nums []int) int {
    result:=int32(0)

	for i:=0;i<32;i++{
		bitCount:=int32(0)
		for _,v:=range nums{
			bitCount+=(int32(v)>>i)&1
		}
		if bitCount%3==1{
			result|=(1<<i)
		}
	}

	return int(result)
}