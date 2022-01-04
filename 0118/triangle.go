package leetcode

func generate(numRows int) [][]int {
	ans := make([][]int, 0, numRows)

	initial := []int{1}
	ans = append(ans, initial)

	for i := 2; i <= numRows; i++ {
		row := make([]int, i)
		row[0], row[i-1] = 1, 1
		for j := 1; j < i-1; j++ {
			row[j] = initial[j-1] + initial[j]
		}
		initial = row
		ans = append(ans, row)
	}
	return ans
}
