package leetcode

func getRow(rowIndex int) []int {
	initial := []int{1}

	for i := 1; i <= rowIndex; i++ {
		row := make([]int, rowIndex+1)
		row[0], row[rowIndex] = 1, 1
		for j := 1; j < i; j++ {
			row[j] = initial[j-1] + initial[j]
		}
		initial = row
	}
	return initial
}
