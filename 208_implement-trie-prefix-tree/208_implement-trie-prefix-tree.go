package leetcode

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

/** Initialize your data structure here. */
func TrieConstructor() Trie {
	res := Trie{}
	return res
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this
	for i := 0; i < len(word); i++ {
		v := word[i] - 'a'
		if cur.children[v] == nil {
			cur.children[v] = &Trie{}
		}
		cur = cur.children[v]
	}
	cur.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this
	for i := 0; i < len(word); i++ {
		v := word[i] - 'a'
		cur = cur.children[v]
		if cur == nil {
			return false
		}
	}
	return cur.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for i := 0; i < len(prefix); i++ {
		v := prefix[i] - 'a'
		cur = cur.children[v]
		if cur == nil {
			return false
		}
	}
	return true
}
