package leetcode

/*
oh my god!  像一堆屎啊！

这道题， 可以用双串的DP 来解决。 是不是很神奇， mark as todo
 */

func isMatch(s string, p string) bool {
	n := len(s)
	m := len(p)

	var visited [][]bool
	var cache [][]bool

	visited = make([][]bool, n)
	cache = make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
		cache[i] = make([]bool, m)
	}

	var dfs func(i, j int) bool

	dfs = func(i, j int) bool {
		//
		if j == m {
			if i == n {
				return true
			}
			return false
		}
		if i == n {
			//need special check
			for k := j; k < m; k++ {
				if p[k] != '*' {
					return false
				}
			}
			return true
		}

		if visited[i][j] {
			return cache[i][j]
		}

		if p[j] == '?' {
			cache[i][j] = dfs(i+1, j+1)
			visited[i][j] = true
			return cache[i][j]
		} else if p[j] == '*' {
			jj := j
			jj++
			distance := 0
			for jj < m && (p[jj] == '*' || p[jj] == '?') {
				if p[jj] == '?' {
					distance++
				}
				jj++
			}
			if jj == m {
				// speical treament for
				/*
					"b"
					"*?*?"
				 */
				if n-i >= distance {
					visited[i][j] = true
					cache[i][j] = true
					return true
				}
				visited[i][j] = true
				cache[i][j] = false
				return false
			}
			stack := make([]int, 0)
			for pos := i; pos < n; pos++ {
				if s[pos] == p[jj] {
					if pos-distance >= i {
						stack = append(stack, pos) // candidate pos
					}
				}
			}

			for k := 0; k < len(stack); k++ {
				if dfs(stack[k]+1, jj+1) {
					cache[i][j] = true
					visited[i][j] = true
					return true
				}
			}
			cache[i][j] = false
			visited[i][j] = true
			return false

		} else {
			if s[i] != p[j] {
				visited[i][j] = true
				cache[i][j] = false
				return false
			}
			cache[i][j] = dfs(i+1, j+1)
			visited[i][j] = true
			return cache[i][j]
		}
	}

	return dfs(0, 0)
}
