package leetcode

type Trie struct {
	children map[byte]*Trie
	complete bool // if complete word
}

func (this *Trie) Add(word string) {
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
	cur.complete = true
}

func dfs(root *Trie, word string, pos int) bool {
	if pos == len(word) {
		return root.complete
	}
	s := word[pos]
	if s != '.' {
		if _, ok := root.children[s]; !ok {
			return false
		}
		return dfs(root.children[s], word, pos+1)
	}

	for _, c := range root.children {
		if dfs(c, word, pos+1) {
			return true
		}
	}
	return false
}

type WordDictionary struct {
	root *Trie
}

func Constructor() WordDictionary {

	return WordDictionary{
		&Trie{
			make(map[byte]*Trie),
			false,
		},
	}
}

func (this *WordDictionary) AddWord(word string) {
	this.root.Add(word)

}

func (this *WordDictionary) Search(word string) bool {
	return dfs(this.root, word, 0)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
