package leetcode

/*
这道题 range dp 是自己想出来的！
 */
func solve(f [][]int, prefixsum []int, n, k int, mergeable []bool) int {

	getter := func(a, b int) int {
		if a > b {
			return 0
		}
		return f[a][b]
	}
	for len := k; len <= n; len++ {
		for i := 0; i < n; i++ {
			j := i + len - 1
			for pos := i; pos <= j && j < n; pos += k - 1 { // pos 每次 increment k-1  这个考虑的比较巧妙啊！ 出 bug 的地方， 因为不是每一种切分都是合理的!
				if mergeable[len] {
					f[i][j] = min(f[i][j], getter(i, pos)+getter(pos+1, j)+prefixsum[j+1]-prefixsum[i])

				} else {
					f[i][j] = min(f[i][j], getter(i, pos)+getter(pos+1, j))
				}
			}
		}
	}
	return f[0][n-1]
}

func min(a, b int) int {
	if a == 0 {
		return b
	}
	if a < b {
		return a
	}
	return b
}

func mergeStones(stones []int, k int) int {
	n := len(stones)
	if n == 1 {
		return 0
	}
	mergeable := make([]bool, n+1)
	for i := k; i <= n; i += k - 1 {
		mergeable[i] = true
	}

	if mergeable[n] == false {
		return -1
	}

	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
	}

	prefixsum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixsum[i] = prefixsum[i-1] + stones[i-1]
	}

	return solve(f, prefixsum, n, k, mergeable)
}
