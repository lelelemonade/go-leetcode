package problem2022

func construct2DArray(original []int, m int, n int) [][]int {
    result:=make([][]int,0)
    if len(original) != m*n{
        return result
    }
    for i:=0;i<m;i++{
        result=append(result,original[i*n:(i+1)*n])
    }
    return result
}