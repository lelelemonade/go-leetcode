package problem120

import "slices"

func minimumTotal(triangle [][]int) int {
	dp:=[]int{triangle[0][0]}
	
	for i:=1;i<len(triangle);i++{
		next:=triangle[i][len(triangle[i])-1]+dp[i-1]

		for j:=len(triangle[i])-2;j>=0;j--{
			if j==0{
				dp[j]=dp[j]+triangle[i][j]
			}else{
				dp[j]=min(dp[j-1],dp[j])+triangle[i][j]
			}
		}
		dp = append(dp, next)
	}
	return slices.Min(dp)
}
