package problem64

func minPathSum(grid [][]int) int {
    dp:=make([]int,len(grid[0]))

	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[0]);j++{
			if i==0{
				if j==0{
					dp[j]=grid[i][j]
				}else{
					dp[j]=dp[j-1]+grid[i][j]
				}
			}else{
				if j==0{
					dp[j]+=grid[i][j]
				}else{
					dp[j]=min(dp[j],dp[j-1])+grid[i][j]
				}
			}
		}
	}
	return dp[len(dp)-1]
}