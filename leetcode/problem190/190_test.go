package problem190

func reverseBits(num uint32) uint32 {
    var result uint32=0

	for i:=0;i<32;i++{
		b:=num%2
		num>>=1

		result<<=1
		result|=b
	}

	return result
}