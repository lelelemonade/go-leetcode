package problem139

//可以用字典树
func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]bool)
	for _, v := range wordDict {
		words[v] = true
	}

	dp := make([]bool, len(s)+1)

	dp[0] = true
	for i := 0; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && words[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
