package leetcode

/*
 这里面的 Index 你搞清楚了吗？
 */
func getModifiedArray(length int, updates [][]int) []int {

	diff := make([]int, length+1)
	// diff[i] = a[i] - a[i-1]  // diff[0] = a[0] - 0
	for i := 0; i < len(updates); i++ {
		l, r, add := updates[i][0], updates[i][1], updates[i][2]
		diff[l] += add
		diff[r+1] -= add
	}

	prefixsum := make([]int, length+1)
	for i := 1; i <= length; i++ {
		prefixsum[i] = prefixsum[i-1] + diff[i-1]
	}
	//ans := make([]int, length)
	//// ans[i] = S[i+1] - S[i]
	//for i := 0; i < length; i++ {
	//	ans[i] = prefixsum[i+1] - prefixsum[i]
	//}
	return prefixsum[1 : length+1]
}
