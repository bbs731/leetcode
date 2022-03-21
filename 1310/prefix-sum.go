package leetcode

// 类似前缀和， 前缀异或
func xorQueries(arr []int, queries [][]int) []int {
	n := len(arr)
	prefix := make([]int, n+1)
	prefix[0] = 0

	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] ^ arr[i-1]
	}

	ans := make([]int, len(queries))

	for i := 0; i < len(queries); i++ {
		query := queries[i]
		ans[i] = prefix[query[1]+1] ^ prefix[query[0]]
	}

	return ans
}
