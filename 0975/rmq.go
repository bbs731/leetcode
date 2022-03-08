package leetcode

import "math"

var maxn [][20]int
var minn [][20]int

func power(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
func RMO_init(arr []int) int {
	n := len(arr)

	maxn = make([][20]int, n)
	minn = make([][20]int, n)

	for i := 0; i < n; i++ {
		maxn[i][0] = arr[i]
		minn[i][0] = arr[i]
	}

	for j := 1; power(2, j) <= n; j++ {
		for i := 0; i+power(2, j)-1 < n; i++ {
			maxn[i][j] = max(maxn[i][j-1], maxn[i+(power(2, j-1))][j-1])
			minn[i][j] = min(minn[i][j-1], minn[i+power(2, j-1)][j-1])
		}
	}
}

func RMQ_max(l, r int) int {
	k := int(math.Log(float64(r-l+1)) * math.Log2E)
	return max(maxn[l][k], maxn[r-power(2, k)+1][k])

}

func RMQ_min(l, r int) int {
	k := int(math.Log(float64(r-l+1)) * math.Log2E)
	return min(minn[l][k], minn[r-power(2, k)+1][k])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
