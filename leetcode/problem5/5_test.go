package problem5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func longestPalindrome(s string) string {
	if len(s)<=1{
		return s
	}

    dp:=make([][]bool,len(s))
	for i:=0;i<len(s);i++{
		dp[i]=make([]bool, len(s))
		dp[i][i]=true
	}
	result:=s[:1]
	for i:=len(s)-2;i>=0;i--{
		for j:=i+1;j<len(s);j++{
			if j-i==1{
				dp[i][j]=s[i]==s[j]
			}else{
				dp[i][j]=s[i]==s[j]&&dp[i+1][j-1]
			}
			if dp[i][j]&&j-i+1>len(result){
				result=s[i:j+1]
			}
		}
	}
	return result
}

func Test5(t *testing.T) {
	assert.Equal(t,"aba",longestPalindrome("babad"))
	assert.Equal(t,"bb",longestPalindrome("cbbd"))
}