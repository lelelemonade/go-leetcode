package problem48

func rotate(matrix [][]int)  {
	n:=len(matrix)
    for i:=0;i<n/2;i++{
		for j:=i;j<n-i-1;j++{
			matrix[i][j],matrix[j][n-1-i],matrix[n-1-i][n-1-j],matrix[n-1-j][i]=
			matrix[n-1-j][i],matrix[i][j],matrix[j][n-1-i],matrix[n-1-i][n-1-j]
		}
	}
}