package problem485

func findMaxConsecutiveOnes(nums []int) int {
	maxConsec:=0
	currentConsec:=0
	for _,v:=range nums{
		if v==1{
			currentConsec++
			maxConsec=max(maxConsec,currentConsec)
		}else{
			currentConsec=0
		}
	}
	return maxConsec
}