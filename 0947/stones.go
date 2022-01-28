package leetcode

/*
不用 DFS ，用 union-find 来找连通量，也是可以的，
但是需要处理 二维到一维的映射，关系. 这个需要技巧，看官方的题解吧！
 */
func removeStones(stones [][]int) int {
	visited := make(map[[2]int]bool)

	rows := make(map[int]map[[2]int]struct{}) // keep all stones at row i
	cols := make(map[int]map[[2]int]struct{}) // keep all stones at col i

	for i := 0; i < len(stones); i++ {
		stone := stones[i]
		row := stone[0]
		col := stone[1]

		if _, ok := rows[row]; !ok {
			rows[row] = make(map[[2]int]struct{})
		}
		rows[row][[2]int{row, col}] = struct{}{}
		if _, ok := cols[col]; !ok {
			cols[col] = make(map[[2]int]struct{})
		}
		cols[col][[2]int{row, col}] = struct{}{}
	}
	// 用 dfs 来找到所有的强连通量
	var dfs func([2]int)
	dfs = func(start [2]int) {
		visited[start] = true

		for k := range rows[start[0]] {
			if visited[k] == false {
				dfs(k)
			}
		}
		for k := range cols[start[1]] {
			if visited[k] == false {
				dfs(k)
			}
		}
	}

	components := 0

	for i := 0; i < len(stones); i++ {
		stone := [2]int{stones[i][0], stones[i][1]}
		if visited[stone] == false {
			components++
			dfs(stone)
		}
	}
	return len(stones) - components
}
