package leetcode

import "strings"

type Trie struct {
	children map[byte]*Trie
	complete bool
}

func (this *Trie) insert(word string) {

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

func (this *Trie) search(word string) string {
	// if found any prefix word return the prefix, otherwise return the word
	cur := this
	for i := 0; i < len(word); i++ {
		s := word[i]
		if _, ok := cur.children[s]; !ok {
			return word
		}
		cur = cur.children[s]
		if cur.complete {
			return word[:i+1]
		}
	}
	return word
}

func replaceWords(dictionary []string, sentence string) string {
	words := strings.Split(sentence, " ")

	// using dictionary to build a trie tree
	root := &Trie{
		make(map[byte]*Trie),
		false,
	}

	for _, item := range dictionary {
		root.insert(item)
	}

	ans := make([]string, 0)
	for _, word := range words {
		ans = append(ans, root.search(word))
	}
	return strings.Join(ans, " ")
}
