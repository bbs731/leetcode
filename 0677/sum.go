package leetcode

type MapSum struct {
	children map[byte]*MapSum
	word     bool
	val      int
}

func Constructor() MapSum {

	sum := MapSum{
		make(map[byte]*MapSum),
		false,
		0,
	}
	return sum
}

func (this *MapSum) Insert(key string, val int) {
	cur := this
	for i := 0; i < len(key); i++ {
		s := key[i]
		if _, ok := cur.children[s]; !ok {
			cur.children[s] = &MapSum{
				make(map[byte]*MapSum),
				false,
				0,
			}
		}
		cur = cur.children[s]
	}
	cur.word = true
	cur.val = val
}

func (this *MapSum) Sum(prefix string) int {
	cur := this
	for i := 0; i < len(prefix); i++ {
		s := prefix[i]
		if _, ok := cur.children[s]; !ok {
			return 0
		}
		cur = cur.children[s]
	}
	// dfs the cur as root
	return dfs(cur)
}

func dfs(root *MapSum) int {

	ans := 0
	if root.word == true {
		ans += root.val
	}
	for _, v := range root.children {
		ans += dfs(v)
	}
	return ans
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
