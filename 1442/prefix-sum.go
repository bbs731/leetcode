package leetcode

func countTriplets(arr []int) int {
	n := len(arr)
	prefix := make([]int, n+1)
	counts := 0

	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] ^ arr[i-1]
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			intern := prefix[j+1] ^ prefix[i]
			if intern == 0 {
				counts += j - i
			}
		}
	}
	return counts
}
