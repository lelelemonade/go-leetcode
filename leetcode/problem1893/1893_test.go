package problem1893

func isCovered(ranges [][]int, left int, right int) bool {

	for _,v:=range ranges{
		if left>=v[0]&&left<=v[1]{
			return true
		}
		if right>=v[0]&&right<=v[1]{
			return true
		}
	}
    
	return false
}