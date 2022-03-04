package leetcode

func minSwap(nums1 []int, nums2 []int) int {

	n := len(nums1)
	//dp[i][0] stands for from 0 to index i, no swap of index i the min swap times
	//dp[i][1] stands fro from 0 to index i, swap of index i the min swap times
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
		dp[i][0] = 100000000
		dp[i][1] = 100000000
	}

	// 初始化
	dp[0][0] = 0
	dp[0][1] = 1

	for i := 1; i < n; i++ {
		// 计算 dp[i][0]
		if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] { // 符合 使用dp[i-1][1] 的条件
			dp[i][0] = min(dp[i][0], dp[i-1][1])
		}
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] { // 符合 使用dp[i-1][0] 的条件
			dp[i][0] = min(dp[i][0], dp[i-1][0])
		}

		// 计算 dp[i][1]
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] { // 符合使用 dp[i-1][1] 的条件
			dp[i][1] = min(dp[i][1], dp[i-1][1]+1)
		}
		if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
			dp[i][1] = min(dp[i][1], dp[i-1][0]+1)
		}
	}

	return min(dp[n-1][0], dp[n-1][1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
