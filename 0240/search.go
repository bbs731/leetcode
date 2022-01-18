package leetcode
/*
我能说什么， 你第一遍写的跟屎一样！
为啥你的反射弧总是那么的长
 */
func binarySearch(nums []int, target int) bool {
	lo := 0
	hi := len(nums) - 1

	for lo <= hi {
		mid := lo + (hi-lo+1)>>1
		if nums[mid] == target {
			return true

		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return false
}

func search(matrix [][]int, row, column, M, N, target int) bool {
	if row > M || column > N {
		return false
	}

	if binarySearch(matrix[row][column:N+1], target) {
		return true
	}
	var c []int
	for i := row; i <= M; i++ {
		c = append(c, matrix[i][column])
	}
	if binarySearch(c, target) {
		return true
	}
	return search(matrix, row+1, column+1, M, N, target)
}

func searchMatrix(matrix [][]int, target int) bool {
	M := len(matrix)
	N := len(matrix[0])
	return search(matrix, 0, 0, M-1, N-1, target)
}
