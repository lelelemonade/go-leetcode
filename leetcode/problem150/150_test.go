package problem150

import "strconv"

func evalRPN(tokens []string) int {
    stack:=make([]int,0)

	for _,v:=range tokens{
		switch v{
		case "+":
			stackLen:=len(stack)
			newValue:=stack[stackLen-2]+stack[stackLen-1]
			stack[stackLen-2]=newValue
			stack=stack[:stackLen-1]
		case "-":
			stackLen:=len(stack)
			newValue:=stack[stackLen-2]-stack[stackLen-1]
			stack[stackLen-2]=newValue
			stack=stack[:stackLen-1]
		case "/":
			stackLen:=len(stack)
			newValue:=stack[stackLen-2]/stack[stackLen-1]
			stack[stackLen-2]=newValue
			stack=stack[:stackLen-1]
		case "*":
			stackLen:=len(stack)
			newValue:=stack[stackLen-2]*stack[stackLen-1]
			stack[stackLen-2]=newValue
			stack=stack[:stackLen-1]
		default:
			stackValue,e:=strconv.Atoi(v)
			if e != nil {
				panic(e)
			}
			stack=append(stack, stackValue)
		}
	}

	return stack[0]
}