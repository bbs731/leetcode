package leetcode

func minSwapsCouples(row []int) int {
	position := make(map[int]int)

	for i := 0; i < len(row); i++ {
		position[row[i]] = i
	}

	N := len(row)
	visited := make([]bool, N)

	var dfs func(int, int) int
	dfs = func(current, target int) int { // return how many nodes have been visited to find target
		visited[current] = true
		pos := position[current]
		var neighbour, npair int

		if pos%2 == 1 {
			neighbour = row[pos-1]
		} else {
			neighbour = row[pos+1]
		}

		visited[neighbour] = true
		if neighbour == target {
			return 2
		}

		// cal neighbour's pair
		if neighbour%2 == 0 {
			npair = neighbour + 1
		} else {
			npair = neighbour - 1
		}
		return 2 + dfs(npair, target)
	}
	ans := 0
	for i := 0; i < N; i++ {
		if visited[i] == false {
			if i%2 == 0 {
				ans += dfs(i, i+1)/2 - 1
			} else {
				ans += dfs(i, i-1)/2 - 1
			}
		}
	}
	return ans
}
