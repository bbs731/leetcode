package leetcode

func setZeroes(matrix [][]int) {

	rows := make(map[int]struct{})
	cols := make(map[int]struct{})

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = struct{}{}
				cols[j] = struct{}{}
			}
		}
	}

	for k := range rows {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[k][j] = 0
		}
	}
	for k := range cols {
		for i := 0; i < len(matrix); i++ {
			matrix[i][k] = 0
		}
	}
}
