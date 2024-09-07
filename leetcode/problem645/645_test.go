package problem645

func findErrorNums(nums []int) []int {
    dup:=0
	total:=0
	set:=make(map[int]struct{})
	for _,v:=range nums{
		if _,e:=set[v];e{
			dup=v
		}
		set[v]=struct{}{}
		total+=v
	}
	return []int{dup,dup+(len(nums)*(len(nums)+1))/2-total}
}