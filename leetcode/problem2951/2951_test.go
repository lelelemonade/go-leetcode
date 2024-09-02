package problem2951

func findPeaks(mountain []int) []int {
    result:=make([]int,0)

	for i:=1;i<len(mountain)-1;i++{
		if mountain[i]>mountain[i-1]&&mountain[i]>mountain[i+1]{
			result = append(result, i)
		}
	}

	return result
}