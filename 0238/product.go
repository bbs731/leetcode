package leetcode

func productExceptSelf(nums []int) []int {
	N := len(nums)

	//prepare two array
	pre := make([]int, N)
	rev := make([]int, N)

	pre[0] = nums[0]
	for i := 1; i < N; i++ {
		pre[i] = nums[i] * pre[i-1]
	}
	rev[N-1] = nums[N-1]
	for j := N - 2; j >= 0; j-- {
		rev[j] = nums[j] * rev[j+1]
	}

	ans := make([]int, N)
	for i := 0; i < N; i++ {
		if i-1 >= 0 {
			ans[i] = pre[i-1]
		} else {
			ans[i] = 1
		}
		if i+1 < N {
			ans[i] *= rev[i+1]
		}
	}
	return ans
}
