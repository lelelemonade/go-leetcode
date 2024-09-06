package problem1832

func checkIfPangram(sentence string) bool {
    var alpha [26]int

	for _,v :=range sentence{
		alpha[v-'a']++
	}
	for _,v:=range alpha{
		if v==0{
			return false
		}
	}
	return true
}