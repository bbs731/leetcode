package leetcode

type Trie struct {
	children map[byte]*Trie
	word     bool
}

func Constructor() Trie {
	t := Trie{
		make(map[byte]*Trie),
		false,
	}
	return t
}

func (this *Trie) Insert(word string) {
	cur := this
	for i := 0; i < len(word); i++ {
		s := word[i]
		if _, ok := cur.children[s]; !ok {
			cur.children[s] = &Trie{
				make(map[byte]*Trie),
				false,
			}

		}
		cur = cur.children[s]
	}
	cur.word = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	for i := 0; i < len(word); i++ {
		s := word[i]
		if _, ok := cur.children[s]; !ok {
			return false
		}
		cur = cur.children[s]
	}
	// since we are search the whole word.
	return cur.word
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for i := 0; i < len(prefix); i++ {
		s := prefix[i]
		if _, ok := cur.children[s]; !ok {
			return false
		}
		cur = cur.children[s]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
