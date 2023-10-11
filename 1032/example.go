package main

// https://cloud.tencent.com/developer/article/1788989

import (
	"fmt"
)

type trieNode struct {
	son     [128]*trieNode // 可以用 128 来处理 ASCII 字符
	fail    *trieNode
	wordend bool
	val     int // pattern's index
}

func (root *trieNode) put(word string, index int) {
	o := root
	for i := range word {
		//idx := word[i] - 'a'
		idx := word[i]
		if o.son[idx] == nil {
			o.son[idx] = &trieNode{}
		}
		o = o.son[idx]
	}
	o.wordend = true
	o.val = index
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

		for i, son := range o.son {
			if son != nil {
				son.fail = o.fail.son[i]
				q = append(q, son)
			} else {
				o.son[i] = o.fail.son[i]
			}
		}
	}
}

func (root *trieNode) match(text string, patterns []string) map[string]int {
	counts := make(map[string]int)

	o := root
	for _, b := range text {
		//idx := b - 'a'
		idx := b
		o = o.son[idx]
		if o == nil {
			o = root
			continue
		}

		for f := o; f != root; f = f.fail {
			if pid := f.val; pid != 0 {
				p := patterns[pid-1]
				counts[p]++
			}
		}
	}
	return counts
}

func main() {
	ac := &trieNode{}
	patterns := []string{"hi", "she", "his", "hers"}

	for i, p := range patterns {
		ac.put(p, i+1)
	}

	ac.buildDFA()

	counts := ac.match("hishihers", patterns)
	for word, count := range counts {
		fmt.Printf("%s: %d\n", word, count)
	}
}
