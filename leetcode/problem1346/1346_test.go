package problem1346

func checkIfExist(arr []int) bool {
    arrMap:=make(map[int]struct{})
	for _,v:=range arr{
		if v==0{
			continue
		}
		if _,e:=arrMap[2*v];e{
			return true
		}
		if _,e:=arrMap[v/2];e&&v%2==0{
			return true
		}
		arrMap[v]=struct{}{}
	}
	return false
}