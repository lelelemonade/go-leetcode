package problem1869

func checkZeroOnes(s string) bool {
	maxConsecOne:=0
	maxConsecZero:=0
	consecOneCount:=0
	consecZeroCount:=0
    for i,v:=range s{
		if v=='0'{
			if i>0&& s[i-1]=='0'{
				consecZeroCount++
			}else{
				consecZeroCount=1
			}
		}else{
			if i>0&&s[i-1]=='1'{
				consecOneCount++
			}else{
				consecOneCount=1
			}
		}
		maxConsecOne=max(maxConsecOne,consecOneCount)
		maxConsecZero=max(maxConsecZero,consecZeroCount)
	}
	return maxConsecOne>maxConsecZero
}