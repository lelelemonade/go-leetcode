package problem1894

func chalkReplacer(chalk []int, k int) int {
    total:=0
	for _,v:=range chalk{
		total+=v
	}
	lastRound:=k%total
	if lastRound<=0{
		return 0
	}
	for i,v:=range chalk{
		lastRound-=v
		if lastRound<0{
			return i
		}else if lastRound==0{
			return i+1
		}
	} 
	return len(chalk)-1
}