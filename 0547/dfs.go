package leetcode

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	visited := make([]bool, n)

	var dfs func(int)
	dfs = func(from int) {
		visited[from] = true

		for i := 0; i < n; i++ {
			if isConnected[from][i] == 1 && visited[i] != true {
				dfs(i)
			}
		}
	}

	components := 0
	for i := 0; i < n; i++ {
		if visited[i] == false {
			components++
			dfs(i)
		}
	}
	return components
}
