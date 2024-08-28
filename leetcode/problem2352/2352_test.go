package problem2352

import "strconv"

func equalPairs(grid [][]int) int {
	rowMap:=make(map[string]int)
	columnMap:=make(map[string]int)

	for i:=0;i<len(grid);i++{
		rowString:=""
		for j:=0;j<len(grid[i]);j++{
			rowString+=strconv.Itoa(grid[i][j])
			rowString+="|"
		}
		rowMap[rowString]++
	}
	for j:=0;j<len(grid[0]);j++{
		colString:=""
		for i:=0;i<len(grid);i++{
			colString+=strconv.Itoa(grid[i][j])
			colString+="|"
		}
		columnMap[colString]++
	}
	result:=0
	for k,rowCount:=range rowMap{
		if columnCount,e:=columnMap[k];e{
			result+=rowCount*columnCount
		}
	}
	return result
}