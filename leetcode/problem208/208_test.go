package problem208

type Trie struct {
	start *TrieNode
}

type TrieNode struct {
	char rune
	end  bool
	next map[rune]*TrieNode
}

func Constructor() Trie {
	return Trie{start: &TrieNode{next: make(map[rune]*TrieNode)}}
}

func (this *Trie) Insert(word string) {
	iter := this.start
	for _, v := range word {
		if n, e := iter.next[v]; e {
			iter = n
		} else {
			nextNode := &TrieNode{char: v, next: make(map[rune]*TrieNode)}
			iter.next[v] = nextNode
			iter = nextNode
		}
	}
	iter.end = true
}

func (this *Trie) Search(word string) bool {
	iter := this.start
	for _, v := range word {
		if n, e := iter.next[v]; !e {
			return false
		} else {
			iter = n
		}
	}
	return iter.end
}

func (this *Trie) StartsWith(prefix string) bool {
	iter := this.start
	for _, v := range prefix {
		if n, e := iter.next[v]; !e {
			return false
		} else {
			iter = n
		}
	}
	return true
}
