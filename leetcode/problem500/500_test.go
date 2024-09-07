package problem500

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

func findWords(words []string) []string {
	row1 := map[byte]struct{}{
		'q': {},
		'w': {},
		'e': {},
		'r': {},
		't': {},
		'y': {},
		'u': {},
		'i': {},
		'o': {},
		'p': {},
	}
	row2 := map[byte]struct{}{
		'a': {},
		's': {},
		'd': {},
		'f': {},
		'g': {},
		'h': {},
		'j': {},
		'k': {},
		'l': {},
	}
	row3 := map[byte]struct{}{
		'z': {},
		'x': {},
		'c': {},
		'v': {},
		'b': {},
		'n': {},
		'm': {},
	}
	result := make([]string, 0)
	check := func(row map[byte]struct{}, word string) {
		wordCount := 0
		for _, v := range word {
			if _, e := row[byte(unicode.ToLower(v))]; e {
				wordCount++
			}
		}
		if wordCount == len(word) {
			result = append(result, word)
		}
	}
	for _, word := range words {
		_, e1 := row1[byte(unicode.ToLower(rune(word[0])))]
		_, e2 := row2[byte(unicode.ToLower(rune(word[0])))]
		_, e3 := row3[byte(unicode.ToLower(rune(word[0])))]
		if e1 {
			check(row1,word)
		} else if e2 {
			check(row2,word)
		} else if e3 {
			check(row3,word)
		}
	}

	return result
}

func Test500(t *testing.T) {
	assert.Equal(t,[]string{"Alaska","Dad"},findWords([]string{"Hello","Alaska","Dad","Peace"}))
}