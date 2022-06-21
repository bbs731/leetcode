package leetcode

func numberWays(hats [][]int) int {
	n := len(hats)
	mod := 1000000000 + 7

	var dfs func(int, uint) int

	dfs = func(level int, bitmask uint) int {
		if level == n {
			return 1
		}
		total := 0

		for _, c := range hats[level] {
			bits := uint(1 << uint(c))
			if bits&bitmask == 0 {
				total += dfs(level+1, bits|bitmask) % mod
				total %= mod
			}
		}
		return total
	}

	return dfs(0, 0)
}
