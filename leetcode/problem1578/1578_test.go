package problem1578

func minCost(colors string, neededTime []int) int {
    result:=0
	i:=0

	for i<len(colors){
		maxTime:=0
		sum:=0
		color:=colors[i]

		for i<len(colors)&&colors[i]==color{
			sum+=neededTime[i]
			maxTime=max(maxTime,neededTime[i])
			i++
		}

		result+=sum-maxTime
	}

	return result
}