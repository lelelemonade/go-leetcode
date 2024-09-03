package problem467

func findSubstringInWraproundString(s string) int {
	dp:=make([]int,26)
	count:=0

	for i:=0;i<len(s);i++{
		if i>0&& (s[i]-s[i-1]+26)%26==1{
			count++
		}else{
			count=1
		}
		dp[s[i]-'a']=max(dp[s[i]-'a'],count)
	}
	result:=0
	for _,v:=range dp{
		result+=v
	}
	return result
}