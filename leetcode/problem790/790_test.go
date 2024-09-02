package problem790

func numTilings(n int) int {
	if n==1{
		return 1
	}
	if n==2{
		return 2
	}
	if n==3{
		return 5
	}
	mod:=1000000007
    dp:=make([]int,n)
	dp[0]=1
	dp[1]=2
	dp[2]=5
	for i:=3;i<n;i++{
		dp[i]=(dp[i-3]+dp[i-1]*2)%mod
	}
	return dp[n-1]
}