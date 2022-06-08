package leetcode

func canFinish(numCourses int, prerequisites [][]int) bool {
	visited := make([]int, numCourses)
	adj := make([][]int, numCourses)

	var dfs func(int) bool
	dfs = func(u int) bool {
		visited[u] = -1
		for _, v := range adj[u] {
			if visited[v] < 0 {
				return false
			} else if visited[v] == 0 {
				if !dfs(v) {
					return false
				}
			}
		}
		visited[u] = 1
		return true
	}

	// prepare adj
	for i := 0; i < numCourses; i++ {
		adj[i] = make([]int, 0)
	}
	for _, p := range prerequisites {
		adj[p[0]] = append(adj[p[0]], p[1])
	}
	// topsort
	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 {
			if !dfs(i) {
				return false
			}
		}
	}
	return true
}
