package problem71

import "strings"

func simplifyPath(path string) string {
    strArray:=strings.Split(path,"/")
	resultStack:=make([]string,0)

	for _,v:=range strArray {
		switch v{
		case ".":
			continue
		case "":
			continue
		case "..":
			if len(resultStack)>0{
				resultStack=resultStack[:len(resultStack)-1]
			}
		default:
			resultStack = append(resultStack, v)
		}
	}
	var resultBuilder strings.Builder

	resultBuilder.WriteString("/")

	for _,v:=range resultStack{
		resultBuilder.WriteString(v)
		resultBuilder.WriteString("/")
	}

	result:=resultBuilder.String()

	if result=="/"{
		return result
	}else{
		return result[:len(result)-1]
	}
}