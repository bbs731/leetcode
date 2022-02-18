package leetcode

func maxKilledEnemies(grid [][]byte) int {

	m := len(grid)
	n := len(grid[0])

	kills := make([][]int, m)
	for i := 0; i < m; i++ {
		kills[i] = make([]int, n)
	}

	// make the border as walls  已达到简化代码的目的
	var getter = func(i, j int) byte {
		if i >= m {
			return 'W'
		}
		if j >= n {
			return 'W'
		}
		return grid[i][j]
	}

	// 先统计行的
	for i := 0; i < m; i++ {
		cnt := 0
		stack := []int{} // save the index
		for j := 0; j <= n; j++ {
			if getter(i, j) == 'E' {
				cnt++
			} else if getter(i, j) == '0' {
				stack = append(stack, j)
			} else {
				// now we have walls. 计算Wall 到上一个 wall 之前的 总共有多少可以杀死的敌人
				for len(stack) > 0 {
					top := stack[len(stack)-1]
					kills[i][top] += cnt
					// pop
					stack = stack[:len(stack)-1]
				}
				//reset
				cnt = 0
			}
		}
	}
	// now treat columns

	for j := 0; j < n; j++ {
		cnt := 0
		stack := []int{}
		for i := 0; i <= m; i++ {
			if getter(i, j) == 'E' {
				cnt++
			} else if getter(i, j) == '0' {
				stack = append(stack, i)
			} else {
				// now we have walls
				for len(stack) > 0 {
					top := stack[len(stack)-1]
					kills[top][j] += cnt
					stack = stack[:len(stack)-1]
				}
				cnt = 0
			}
		}
	}

	ans := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				ans = max(ans, kills[i][j])
			}
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
