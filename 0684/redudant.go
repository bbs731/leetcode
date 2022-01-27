package leetcode

/*
union-find 的解法！
 */
/*
union find 有标准的实现方法： find() 函数， 和 union 函数. 写一个标准的模板！
 */
func findRedundantConnection(edges [][]int) []int {

	parent := make(map[int]int)
	var find func(int) int
	find = func(n int) int { // if n doest no have parent return n itself, otherwise return its root
		if _, ok := parent[n]; !ok {
			parent[n] = n
		}
		if parent[n] != n {
			parent[n] = find(parent[n])
		}
		return parent[n]
	}

	for i := 0; i < len(edges); i++ {
		pairs := edges[i]

		p1 := find(pairs[0])
		p2 := find(pairs[1])

		if p1 == p2 {
			return pairs
		}
		// union
		parent[p1] = p2
	}

	return []int{}
}
