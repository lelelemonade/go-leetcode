package problem20

func isValid(s string) bool {
    stack := make([]rune,0)

	for _,v:=range s{
		switch v {
		case '{':
			stack = append(stack, v)
		case '[':
			stack = append(stack, v)
		case '(':
			stack = append(stack, v)
		case '}':
			if len(stack)!=0 && stack[len(stack)-1] == '{'{
				stack=stack[:len(stack)-1]
			}else{
				return false
			}
		case ']':
			if len(stack)!=0 && stack[len(stack)-1] == '['{
				stack=stack[:len(stack)-1]
			}else{
				return false
			}
		case ')':
			if len(stack)!=0 && stack[len(stack)-1] == '('{
				stack=stack[:len(stack)-1]
			}else{
				return false
			}
		}
	}
	return len(stack)==0
}