package problem67

import "slices"

func addBinary(a string, b string) string {
    result:=make([]byte,0)
	carry:=0

	for ai,bi:=len(a)-1,len(b)-1;ai>=0||bi>=0;{
		var aBit byte
		if ai<0{
			aBit='0'
		}else{
			aBit=a[ai]
		}

		var bBit byte
		if bi<0{
			bBit='0'
		}else{
			bBit=b[bi]
		}

		if aBit=='1'&&bBit=='0'{
			if carry==1{
				result=append(result, '0')
			}else{
				result=append(result, '1')
			}
		}else if aBit=='0'&&bBit=='1'{
			if carry==1{
				result=append(result, '0')
			}else{
				result=append(result, '1')
			}
		}else if aBit=='1'&&bBit=='1'{
			if carry==1{
				result=append(result, '1')
			}else{
				result=append(result, '0')
				carry=1
			}
		}else{
			if carry==1{
				result=append(result, '1')
				carry=0
			}else{
				result=append(result, '0')
			}
		}

		ai--
		bi--
	}
	if carry==1{
		result = append(result, '1')
	}
	slices.Reverse(result)
	return string(result)
}