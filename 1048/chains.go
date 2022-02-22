package leetcode

import "sort"

/*
dfs 的解法。

还可以用 DP 的方法！ try that !
 */
func longestStrChain(words []string) int {

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	grids := make([][]bool, len(words))
	for i := 0; i < len(words); i++ {
		grids[i] = make([]bool, len(words))
		for j := i + 1; j < len(words); j++ {

			if len(words[i]) == len(words[j]) {
				continue
			} else if len(words[i])+1 == len(words[j]) {
				if diffbyone(words[i], words[j]) {
					grids[i][j] = true
				}
			} else {
				// len(words[i]) +1  < len(words[j]
				break
			}
		}
	}
	// now we have connected graph grids. let's find the longest path

	visited := make([]bool, len(words))
	f := make([]int, len(words))

	for i := 0; i < len(words); i++ {
		f[i] = 1
	}

	var dfs func(n int) int
	dfs = func(n int) int {
		if visited[n] == true {
			return f[n]
		}

		// check all neighbours
		for i := n + 1; i < len(words); i++ {
			if grids[n][i] {
				f[n] = max(f[n], dfs(i)+1)
			}
		}
		visited[n] = true
		return f[n]
	}
	ans := 1
	for i := 0; i < len(words); i++ {
		if visited[i] == false {
			ans = max(ans, dfs(i))
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func diffbyone(s, t string) bool {
	diff := false // only allow one diff
	for i, j := 0, 0; i < len(s); {
		if s[i] == t[j] {
			j++
			i++
			continue
		} else {
			if diff {
				return false
			} else {
				j++
				diff = true
			}
		}
	}
	return true
}
