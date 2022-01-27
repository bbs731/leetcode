package leetcode

/*
union-find 的解法！ 
 */
func findRedundantConnection(edges [][]int) []int {

	parent := make(map[int]int)
	findRoot := func(n int) int { // if n doest no have parent return n itself, otherwise return its root
		var p int
		if _, ok := parent[n]; !ok {
			parent[n] = n
		}
		p = parent[n]
		for p != n {
			n = p
			p = parent[n]
		}
		return p
	}

	for i := 0; i < len(edges); i++ {
		pairs := edges[i]

		p1 := findRoot(pairs[0])
		p2 := findRoot(pairs[1])

		if p1 == p2 {
			return pairs
		}
		// join
		parent[p1] = p2
	}

	return []int{}
}
