package leetcode

/*
Trie 的代码实现：

代码源头：
https://github.com/EndlessCheng/codeforces-go/blob/master/copypasta/trie.go

https://leetcode.cn/problems/implement-trie-ii-prefix-tree/solution/go-xie-fa-by-endlesscheng-i16o/
 */

type trieNode struct {
	son    [26]*trieNode
	prefix int
	end    int
}

func (o *trieNode) empty() bool {
	for _, son := range o.son {
		if son != nil {
			return false
		}
	}
	return true
}

type Trie struct {
	root *trieNode
}

func (Trie) ord(c byte) byte { return c - 'a' }
func (Trie) chr(v byte) byte { return v + 'a' }

func Constructor() Trie {

	return Trie{&trieNode{}}
}

func (this *Trie) Insert(word string) {
	o := this.root

	for _, b := range []byte(word) {
		b = this.ord(b)
		if o.son[b] == nil {
			o.son[b] = &trieNode{}
		}
		o = o.son[b]
		o.prefix++
	}
	o.end++
}

func (this *Trie) CountWordsEqualTo(word string) int {
	o := this.root
	for _, b := range []byte(word) {
		o = o.son[this.ord(b)]
		if o == nil {
			return 0
		}
	}
	return o.end
}

func (this *Trie) CountWordsStartingWith(prefix string) int {
	o := this.root

	for _, b := range prefix {
		o = o.son[this.ord(byte(b))]
		if o == nil {
			return 0
		}
	}
	return o.prefix
}

func (this *Trie) Erase(word string) {
	o := this.root
	for _, b := range word {

		o = o.son[this.ord(byte(b))]
		if o == nil {
			return
		}
		o.prefix--
	}
	o.end--

}
