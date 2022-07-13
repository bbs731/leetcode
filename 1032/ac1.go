package leetcode

type trieNode struct {
	son     [26]*trieNode
	fail    *trieNode
	wordend bool
}

func (root *trieNode) put(word string) {
	o := root
	for i := range word {
		idx := word[i] - 'a'
		if o.son[idx] == nil {
			o.son[idx] = &trieNode{}
		}
		o = o.son[idx]
	}
	o.wordend = true
}

func (root *trieNode) buildDFA() {
	root.fail = root
	q := []*trieNode{}

	for _, son := range root.son[:] {
		if son != nil {
			son.fail = root
			q = append(q, son)
		}
	}

	for len(q) > 0 {
		o := q[0]
		q = q[1:]
		if o.fail == nil {
			o.fail = root
		}
		//o.wordend = o.wordend || o.fail.wordend // 根据题目的解， 简化的一种写法, 也可以把这个逻辑放在 query 里。
		for i, son := range o.son[:] {
			if son != nil {
				son.fail = o.fail.son[i]
				q = append(q, son)
			} else {
				o.son[i] = o.fail.son[i]
			}
		}
	}
}

type StreamChecker struct {
	root *trieNode
	cur  *trieNode
}

func Constructor(words []string) StreamChecker {
	root := &trieNode{}

	for _, w := range words {
		root.put(w)
	}
	root.buildDFA()
	return StreamChecker{root, root}
}

func (this *StreamChecker) Query(letter byte) bool {
	this.cur = this.cur.son[letter-'a']
	if this.cur == nil {
		this.cur = this.root
	}
	ret := this.cur.wordend

	for f := this.cur.fail; f != this.root; f = f.fail {
		ret = ret || f.wordend
	}
	return ret
}
