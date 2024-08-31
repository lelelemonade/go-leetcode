package problem1268

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TrieNode struct {
	children map[rune]*TrieNode
	words    []string
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
		node.words = append(node.words, word)
		// 只保留前三个字典序的词
		sort.Strings(node.words)
		if len(node.words) > 3 {
			node.words = node.words[:3]
		}
	}
}

func (t *Trie) Search(prefix string) []string {
	node := t.root
	for _, ch := range prefix {
		if _, ok := node.children[ch]; !ok {
			return []string{}
		}
		node = node.children[ch]
	}
	return node.words
}

func suggestedProducts(products []string, searchWord string) [][]string {
	trie := Constructor()

	for _, product := range products {
		trie.Insert(product)
	}

	result := [][]string{}
	prefix := ""
	for _, ch := range searchWord {
		prefix += string(ch)
		result = append(result, trie.Search(prefix))
	}

	return result
}

func Test1268(t *testing.T) {

	assert.Equal(t, [][]string{}, suggestedProducts([]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse"))
}
