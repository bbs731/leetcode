package leetcode

// ToDo: trailing recursion 可以写成 loop 形式。 锻炼一下?
func traverse(matrix [][]int, m int, n int, origin int) []int {
	trace := make([]int, 0)
	if m == 0 || n == 0 {
		return []int{}
	}
	if m == 1 {
		// take the only row
		for i := origin; i < origin+n; i++ {
			trace = append(trace, matrix[origin][i])
		}
		return trace
	}

	if n == 1 {
		// take the only column
		for j := origin; j < origin+m; j++ {
			trace = append(trace, matrix[j][origin])
		}
		return trace
	}

	//otherwise we clockwise traverse the matrix
	for i := origin; i < origin+n; i++ {
		trace = append(trace, matrix[origin][i])
	}
	//take the right column
	for j := origin + 1; j < origin+m; j++ {
		trace = append(trace, matrix[j][origin+n-1])
	}
	for i := origin + n - 2; i >= origin; i-- {
		trace = append(trace, matrix[origin+m-1][i])
	}
	for j := origin + m - 2; j > origin; j-- {
		trace = append(trace, matrix[j][origin])
	}

	trace = append(trace, traverse(matrix, m-2, n-2, origin+1)...)
	return trace
}

func spiralOrder(matrix [][]int) []int {
	// get the matrix row and colum numbers
	m := len(matrix)
	n := len(matrix[0])

	return traverse(matrix, m, n, 0)
}
