package problem221

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maximalSquare(matrix [][]byte) int {
    dp:=make([][]int,len(matrix))

	for i:=0;i<len(matrix);i++{
		dp[i]=make([]int, len(matrix[0]))
	}

	maxSquare:=0

	for i:=0;i<len(matrix);i++{
		for j:=0;j<len(matrix[0]);j++{
			if i==0||j==0{
				if matrix[i][j]=='1'{
					dp[i][j]=1
				}else{
					dp[i][j]=0
				}
			}else{
				if matrix[i][j]=='1'{
					dp[i][j]=min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1])+1
				}else{
					dp[i][j]=0
				}
			}
			maxSquare=max(maxSquare,dp[i][j]*dp[i][j])
		}
	}
	return maxSquare
}

func Test221(t *testing.T) {
	// assert.Equal(t,4,maximalSquare([][]byte{{'1','0','1','0','0'},{'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}}))
	assert.Equal(t,1,maximalSquare([][]byte{{'0','1'}}))
}