package problem1732

func largestAltitude(gain []int) int {
    altitude:=0
	result:=0

	for _,v:=range gain{
		altitude+=v
		result=max(result,altitude)
	}

	return result
}