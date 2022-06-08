package leetcode

func findOrder(numCourses int, prerequisites [][]int) []int {
	var dfs func(int) bool
	visited := make([]int, numCourses)
	orders := make([]int, 0, numCourses)
	adj := make([][]int, numCourses)

	dfs = func(u int) bool {
		visited[u] = -1
		for _, v := range adj[u] {
			if visited[v] == -1 {
				return false
			} else if visited[v] == 0 {
				if !dfs(v) {
					return false
				}
			}
		}
		visited[u] = 1
		orders = append(orders, u)
		return true
	}

	// prepare adj
	for _, p := range prerequisites {
		adj[p[0]] = append(adj[p[0]], p[1])
	}

	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 {
			if !dfs(i) {
				return []int{}
			}
		}
	}
	return orders
}
