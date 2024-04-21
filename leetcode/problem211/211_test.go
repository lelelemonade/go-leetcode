package problem211

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test211(t *testing.T) {
	wordDictionary := Constructor()

	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")

	assert.False(t, wordDictionary.Search("pad"))
	assert.True(t, wordDictionary.Search("bad"))
	assert.True(t, wordDictionary.Search(".ad"))
	assert.True(t, wordDictionary.Search("b.."))
}

type TrieNode struct {
	child map[byte]*TrieNode
	isEnd bool
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{root: &TrieNode{}}

}

func (this *WordDictionary) AddWord(word string) {
	iter := this.root

	for i := 0; i < len(word); i++ {
		char := word[i]

		if iter.child == nil {
			iter.child = make(map[byte]*TrieNode)
		}
		if iter.child[char] == nil {
			iter.child[char] = &TrieNode{}
		}

		iter = iter.child[char]
	}

	iter.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(int, *TrieNode) bool

	dfs = func(i int, iter *TrieNode) bool {
		if i == len(word) {
			return iter.isEnd
		}

		if word[i] == '.' {
			result := false
			for _, v := range iter.child {
				result = result || dfs(i+1, v)
			}

			return result
		}

		if iter.child[word[i]] == nil {
			return false
		}

		return dfs(i+1, iter.child[word[i]])
	}

	return dfs(0, this.root)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
