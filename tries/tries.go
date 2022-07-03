package tries

// Implement Trie (Prefix Tree)
// https://leetcode.com/problems/implement-trie-prefix-tree/
//

type TrieNode struct {
	is_end bool
	keys   map[rune]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		is_end: false,
		keys:   map[rune]*TrieNode{},
	}
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: NewTrieNode(),
	}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for _, c := range word {
		if _, ok := node.keys[c]; !ok {
			t := NewTrieNode()
			node.keys[c] = t
		}
		node = node.keys[c]
	}
	node.is_end = true
}

func (this *Trie) Search(word string) bool {
	result, node := false, this.root
	for _, c := range word {
		if n, ok := node.keys[c]; ok {
			result = n.is_end
			node = n
		} else {
			return false
		}
	}
	return result
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, c := range prefix {
		if n, ok := node.keys[c]; ok {
			node = n
		} else {
			return false
		}
	}
	return true
}
