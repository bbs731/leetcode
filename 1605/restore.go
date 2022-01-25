package leetcode

import "sort"

type Pair struct {
	val   int
	index int
}

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	ans := make([][]int, 0)

	// fill row by row
	for i := 0; i < len(rowSum); i++ {
		// distribute rowSum[i] among all columns according to current colSum's order

		cols := make([]Pair, len(colSum))
		for j := 0; j < len(colSum); j++ {
			cols[j] = Pair{colSum[j], j}
		}

		sort.Slice(cols, func(i, j int) bool { return cols[i].val > cols[j].val }) // largest first

		sum := rowSum[i]
		row := make([]int, len(colSum))
		for j := 0; j < len(colSum); j++ {
			index := cols[j].index
			if sum <= colSum[index] {
				row[index] = sum
				ans = append(ans, row)
				// update cols
				colSum[index] -= sum
				break
			} else {
				row[index] = colSum[index]
				sum -= colSum[index]
				colSum[index] = 0
			}
		}
	}
	return ans
}
