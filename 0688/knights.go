package leetcode

import (
	"fmt"
	"math"
)

/*
这道题，给我们的教训就是， 指数级的数，非常的大， 需要用 float64, int64会 overflow, 另外。 不能用 stack 来保持中间的状态，
状态太多以至于 out of memory
 */
func knightProbability(n int, k int, row int, column int) float64 {

	directions := [][]int{[]int{1, 2}, []int{2, 1}, []int{2, -1}, []int{1, -2}, []int{-1, -2}, []int{-2, -1}, []int{-2, 1}, []int{-1, 2}}

	dp := make([][]float64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]float64, n)
	}

	check := func(pos [2]int) bool {
		if pos[0] < 0 || pos[0] > n-1 || pos[1] < 0 || pos[1] > n-1 {
			return false
		}
		return true
	}
	move := func(pos [2]int, res [][]float64) {
		for j := 0; j < len(directions); j++ {
			newmove := [2]int{pos[0] + directions[j][0], pos[1] + directions[j][1]}
			if check(newmove) {
				res[newmove[0]][newmove[1]] += dp[pos[0]][pos[1]]
			}
		}
	}

	dp[row][column] = 1

	cnt := 1
	for cnt <= k {
		stack := make([][2]int, 0)

		next := make([][]float64, n)
		for i := 0; i < n; i++ {
			next[i] = make([]float64, n)
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dp[i][j] != 0 {
					stack = append(stack, [2]int{i, j})
				}
			}
		}

		for len(stack) != 0 {
			top := stack[len(stack)-1]
			// give the move
			move(top, next)
			stack = stack[:len(stack)-1]
		}
		dp = next
		cnt++
		fmt.Println(dp)
	}

	var valids float64
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] != 0 {
				valids += dp[i][j]
			}
		}
	}

	return float64(valids) / math.Pow(8, float64(k))
}
