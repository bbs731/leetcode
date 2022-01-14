package leetcode

import "math"

func superEggDrop(k int, n int) int {

	// dp[n][k] 代表的是， n 层楼， k 个 eggs, 最少能得到 f 的次数
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}

	// 初始化
	for eggs := 1; eggs <= k; eggs++ {
		dp[1][eggs] = 1
	}

	for i := 1; i <= n; i++ {
		dp[i][1] = i
	}

	for i := 2; i <= n; i++ {
		// eggs 从 2 开始， eggs 等于 1 是 dp[i][1] 是在初始化中完成的， 在 dp[j-1][eggs-1]
		for eggs := 2; eggs <= k; eggs++ {
			dp[i][eggs] = math.MaxInt64

			lo, hi := 1, i

			/*
				eggs no break := dp[i-j][eggs] + 1
				// otherwise jth floor, eggs break
				eggs breaks :=  dp[j-1][eggs-1] + 1

			在 eggs 给定， i 给定的情况下， j 是变量。  那么， dp[i-j][eggs] 是关于 j 的递减函数
			dp[j-1][eggs-1] 是关于 j 变量的递增函数。 因此，可以用二分查找的方式， 来找到最优的 j 值, 即 函数 dp[i-j][eggs] == dp[j-1][eggs-1]
			具体题解，参看， leetcode 887 的官方题解。

			 */

			for lo+1 < hi {
				mid := lo + (hi-lo+1)/2
				// suppose jth floor , eggs no break.
				t2 := dp[i-mid][eggs]
				// otherwise jth floor, eggs break
				t1 := dp[mid-1][eggs-1]

				if t1 < t2 {
					lo = mid
				} else if t1 > t2 {
					hi = mid
				} else {
					lo = mid
					hi = mid
				}
			}
			dp[i][eggs] = min(max(dp[lo-1][eggs-1], dp[i-lo][eggs]), max(dp[hi-1][eggs-1], dp[i-hi][eggs]))
		}
	}
	return dp[n][k]
}
