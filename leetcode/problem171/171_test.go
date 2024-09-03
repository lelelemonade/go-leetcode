package problem171

func titleToNumber(columnTitle string) int {
    result:=0
	digit:=1
	for i:=len(columnTitle)-1;i>=0;i--{
		result+=int(columnTitle[i]-'A'+1)*digit
		digit*=26
	}
	return result
}