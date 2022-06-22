package leetcode

/*
typical bitmask DP problem,
w/o any memoryization,  n, m size is 10
 */
func assignBikes(workers [][]int, bikes [][]int) int {

	n := len(workers)
	m := len(bikes)
	var dfs func(int, uint, int)

	ans := 1000000 * 3

	// we mask the bike
	dfs = func(worker int, mask uint, cost int) {
		if worker == n {
			ans = min(ans, cost)
			return
		}
		for i := 0; i < m; i++ {
			bits := uint(1 << uint(i))
			if mask&bits == 0 { //bike not used
				dfs(worker+1, mask|bits, cost+distance(workers[worker], bikes[i]))
			}
		}
	}
	dfs(0, 0, 0)
	return ans
}

func distance(a, b []int) int {
	return abs(a[0], b[0]) + abs(a[1], b[1])
}

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
