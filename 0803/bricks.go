package leetcode

func hitBricks(grid [][]int, hits [][]int) []int {
	m := len(grid)
	n := len(grid[0])

	var neigh func(int, int) [][2]int
	neigh = func(i, j int) [][2]int {
		valid := func(i, j int) bool {
			if i >= 0 && i < m && j >= 0 && j < n {
				return true
			}
			return false
		}
		l := [][2]int{[2]int{i - 1, j}, [2]int{i, j + 1}, [2]int{i, j - 1}, [2]int{i + 1, j}}
		res := make([][2]int, 0)
		for _, p := range l {
			if valid(p[0], p[1]) && grid[p[0]][p[1]] == 1 {
				res = append(res, p)
			}
		}
		return res
	}

	ans := make([]int, 0)
	// 用 dfs 来找到 grid 中的强连通量

	for i := 0; i < len(hits); i++ {
		hit := hits[i]
		if grid[hit[0]][hit[1]] == 0 {
			ans = append(ans, 0)
			continue
		}
		grid[hit[0]][hit[1]] = 0

		visited := make(map[[2]int]struct{})

		var dfs func(int, int) (bool, [][2]int)
		dfs = func(i, j int) (stable bool, connected [][2]int) {
			// mark i, j visited
			visited[[2]int{i, j}] = struct{}{}
			found := false // during dfs want to find some brick has row as 0 , which is at the top so stable
			if i == 0 { // top brick at the ceiling
				found = true
			}
			saved := make([][2]int, 0)

			for _, n := range neigh(i, j) {
				if _, ok := visited[n]; !ok {
					f, l := dfs(n[0], n[1])
					if f == true {
						found = true
						return found, [][2]int{}
					}
					saved = append(saved, l...)
				}
			}

			saved = append(saved, [2]int{i, j})
			return found, saved
		}

		total := 0
		toremove := make([][2]int, 0)
		for _, neighbour := range neigh(hit[0], hit[1]) {
			if _, ok := visited[[2]int{neighbour[0], neighbour[1]}]; !ok { // if not visited before
				found, l := dfs(neighbour[0], neighbour[1])
				if !found {
					total += len(l)
					toremove = append(toremove, l...)
				}
			}
		}
		ans = append(ans, total)
		// now we need to remove "toremove bricks"
		for _, r := range toremove {
			grid[r[0]][r[1]] = 0
		}
	}

	return ans
}
