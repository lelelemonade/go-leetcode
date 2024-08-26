package problem1431

func kidsWithCandies(candies []int, extraCandies int) []bool {

    maxCandy:=0
	for _,v:=range candies{
		maxCandy=max(maxCandy,v)
	}
	result:=make([]bool,len(candies))
	for i:=0;i<len(result);i++{
		if candies[i]+extraCandies>=maxCandy{
			result[i]=true
		}
	}
	return result
}