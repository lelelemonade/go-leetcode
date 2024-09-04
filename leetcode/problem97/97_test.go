package problem97

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3)!=len(s1)+len(s2){
		return false
	}
	dp:=make([][]bool,len(s1)+1)
	for i:=0;i<len(dp);i++{
		dp[i]=make([]bool,len(s2)+1)
	}

	dp[0][0]=true

	for i:=0;i<len(dp);i++{
		for j:=0;j<len(dp[0]);j++{
			p := i + j - 1
            if i > 0 {
                dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[p])
            }
            if j > 0 {
                dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[p])
            }
		}
	}

	return dp[len(dp)-1][len(dp[0])-1]
}

func Test97(t *testing.T) {
	assert.Equal(t,false,isInterleave("aabcc","dbbca","aadbbcccac"))
	assert.Equal(t,true,isInterleave("aabc","abad","aabadabc"))
}