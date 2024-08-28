package problem2390

func removeStars(s string) string {
    stack := make([]rune,0)

	for _,v:=range s{
		if v=='*'{
			if len(stack)>0{
				stack=stack[:len(stack)-1]
			}
		}else{
			stack=append(stack, v)
		}
	}

	return string(stack)
}