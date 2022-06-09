package leetcode

/*
这道题，涉及到的一个核心问题是： 如何找到一个树中的最长路径？

路径中间的 点（或者两个点） 就是这道题的答案。
 */
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	g := make([][]int, n)

	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}

	parents := make([]int, n)
	visited := make([]bool, n)
	path := make([]int, 0)
	maxdepth := 0
	farnode := -1

	for i := 0; i < n; i++ {
		parents[i] = -1
		visited[i] = false
	}
	var dfs func(int, int)
	dfs = func(s int, level int) {
		visited[s] = true
		if level > maxdepth {
			maxdepth = level
			farnode = s
		}
		for _, v := range g[s] {
			if visited[v] == false {
				parents[v] = s
				dfs(v, level+1)
			}
		}
	}
	dfs(0, 0)
	x := farnode
	farnode = -1
	maxdepth = 0
	for i := 0; i < n; i++ {
		parents[i] = -1
		visited[i] = false
	}

	dfs(x, 0)
	// collect path
	for farnode != -1 {
		path = append(path, farnode)
		farnode = parents[farnode]
	}

	if len(path)%2 == 0 {
		l := len(path)
		return []int{path[l/2-1], path[l/2]}

	} else {
		return []int{path[len(path)/2]}
	}
}
