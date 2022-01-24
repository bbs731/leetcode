package leetcode
/*
如果，做剪枝，优化呢？

算了吧， 我觉得，这套题的解题思路还是挺清晰的！
 */
func search(board [][]byte, visited [][]bool, word string, k, row, col int) bool {

	if k == len(word) {
		// found
		return true
	}

	m := len(board)
	n := len(board[0])

	getter := func(i, j int) bool {
		if i < 0 || i >= m {
			return true
		}
		if j < 0 || j >= n {
			return true
		}
		return visited[i][j]
	}

	getboard := func(i, j int) byte {
		if i < 0 || i >= m {
			return '$'
		}
		if j < 0 || j >= n {
			return '$'
		}
		return board[i][j]
	}

	pos := [][]int{[]int{row - 1, col}, []int{row + 1, col}, []int{row, col + 1}, []int{row, col - 1}}
	for _, v := range pos {
		r := v[0]
		c := v[1]
		if getter(r, c) == false && getboard(r, c) == word[k] { // not visited before and match
			visited[r][c] = true
			if search(board, visited, word, k+1, r, c) {
				return true
			}
			visited[r][c] = false
		}
	}
	return false
}

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				visited[i][j] = true
				if search(board, visited, word, 1, i, j) {
					return true
				}
				visited[i][j] = false
			}
		}
	}
	return false
}
