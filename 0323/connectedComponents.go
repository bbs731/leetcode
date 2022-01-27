package leetcode

func countComponents(n int, edges [][]int) int {

	visited := make(map[int]bool, n)
	components := 0

	tables := make(map[int]map[int]struct{}, n)
	for i := 0; i < n; i++ {
		tables[i] = make(map[int]struct{})
	}

	for i := 0; i < len(edges); i++ {
		pair := edges[i]
		// 无向图
		tables[pair[0]][pair[1]] = struct{}{}
		tables[pair[1]][pair[0]] = struct{}{}
	}

	var dfs func(int)
	dfs = func(node int) {
		if visited[node] == true {
			return
		}
		visited[node] = true
		for k := range tables[node] {
			dfs(k)
		}
	}

	for i := 0; i < n; i++ {
		if _, ok := visited[i]; ok {
			continue
		}
		dfs(i)
		components++
	}
	return components
}
