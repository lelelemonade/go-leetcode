package problem516

func longestPalindromeSubseq(s string) int {
	dp:=make([][]int,len(s))
	for i:=0;i<len(s);i++{
		dp[i]=make([]int,len(s))
		dp[i][i]=1
	}

	for i:=len(s)-2;i>=0;i--{
		for j:=i+1;j<len(s);j++{
			if s[i]==s[j]{
				dp[i][j]=dp[i+1][j-1]+2
			}else{
				dp[i][j]=max(dp[i+1][j],dp[i][j-1])
			}
		}
	}
	return dp[0][len(s)-1]
}