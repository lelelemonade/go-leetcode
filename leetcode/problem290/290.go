package problem290

import "strings"

func wordPattern(pattern string, s string) bool {
	ch2word := make(map[byte]string)
	word2ch := make(map[string]byte)
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}
	for i, word := range words {
		ch := pattern[i]
		if word2ch[word] > 0 && word2ch[word] != ch || ch2word[ch] != "" && ch2word[ch] != word {
			return false
		}
		word2ch[word] = ch
		ch2word[ch] = word
	}
	return true
}
