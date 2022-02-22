package leetcode

/*
this is a backtrace problem ,not  a DP problem!
 */
func getFactors(n int) [][]int {

	ans := make([][]int, 0)

	var dfs func(int, []int)
	dfs = func(l int, path []int) {
		if l == 1 {
			// accumulate the answer
			// copy the path and added to ans
			var candidate []int
			candidate = append(candidate, path...)
			ans = append(ans, candidate)
			return
		}

		for i := path[len(path)-1]; i*i <= l; i++ {
			if l%i == 0 {
				path = append(path, i)
				dfs(l/i, path)
				path = path[:len(path)-1]
			}
		}
		path = append(path, l)
		dfs(1, path)
		path = path[:len(path)-1]
	}

	path := make([]int, 0)

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			path = append(path, i)
			dfs(n/i, path)
			// restore
			path = path[:len(path)-1]
		}
	}
	return ans
}
