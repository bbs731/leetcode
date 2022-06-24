package leetcode

/*
这个 trie 的实现，很蹩脚啊！ 看看其它例子的Trie实现
 */
type Trie struct {
	word bool
	s    string
	next map[byte]*Trie
}

func (t *Trie) Insert(w string) {
	n := len(w)
	r := t
	for i := 0; i < n; i++ {
		if r.next[w[i]] == nil {
			r.next[w[i]] = &Trie{
				false,
				"",
				make(map[byte]*Trie),
			}
		}
		r = r.next[w[i]]
		if i == n-1 {
			r.word = true
			r.s = w
		}
	}
}

func (t *Trie) Traverse(s string) []string {
	ans := make([]string, 0)
	r := t
	for i := 0; i < len(s); i++ {
		if r == nil {
			break
		}
		if r.word {
			ans = append(ans, r.s)
		}
		r = r.next[s[i]]
	}
	if r != nil && r.word { // 这里，让你困住了好长时间！
		ans = append(ans, r.s)
	}

	return ans
}

func wordBreak(s string, wordDict []string) []string {
	root := &Trie{
		false,
		"",
		make(map[byte]*Trie),
	}

	for _, w := range wordDict {
		root.Insert(w)
	}

	ans := make([]string, 0)
	var dfs func(string, string)

	dfs = func(s string, sentence string) {

		if s == "" {
			ans = append(ans, sentence[1:])
			return
		}

		candidates := root.Traverse(s)
		for _, c := range candidates {
			dfs(s[len(c):], sentence+" "+c)
		}
	}
	dfs(s, "")

	return ans
}
