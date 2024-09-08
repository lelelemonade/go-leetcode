package problem91

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func numDecodings(s string) int {
    dp:=make([]int,len(s)+1)
	dp[0]=1
	if s[0]=='0'{
		dp[1]=0
	}else{
		dp[1]=1
	}
	
	for i:=2;i<len(s)+1;i++{
		num:=s[i-1]
		if num>='1'&&num<='6'{
			if s[i-2]=='1'||s[i-2]=='2'{
				dp[i]=dp[i-1]+dp[i-2]
			}else{
				dp[i]=dp[i-1]
			}
		}else if num=='0' {
			if s[i-2]=='1'||s[i-2]=='2'{
				dp[i]=dp[i-2]
			}else{
				dp[i]=0
			}
		}else{
			if s[i-2]=='1'{
				dp[i]=dp[i-1]+dp[i-2]
			}else{
				dp[i]=dp[i-1]
			}
		}
	}
	return dp[len(dp)-1]
}

func Test91(t *testing.T) {
	// assert.Equal(t,2,numDecodings("12"))
	// assert.Equal(t,3,numDecodings("226"))
	assert.Equal(t,0,numDecodings("06"))
}