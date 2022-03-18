package leetcode

/*
这个版本的， index 处理的优雅些！
 */
func productExceptSelf(nums []int) []int {
	n := len(nums)

	ans := make([]int, n)

	//prepare two array  前缀积， 和后缀积
	prefix := make([]int, n+1)
	reverse := make([]int, n+1)

	prefix[0] = 1
	reverse[n] = 1
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}
	for i := n - 1; i >= 0; i-- {
		reverse[i] = reverse[i+1] * nums[i]
	}

	for i := 0; i < n; i++ {
		ans[i] = prefix[i] * reverse[i+1]
	}
	return ans
}
