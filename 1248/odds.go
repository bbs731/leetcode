package leetcode

/*
还不是很理解，题的意思， 一次提交就过了！ 哎！ 欲速而不达！
 */

func numberOfSubarrays(nums []int, k int) int {

	ans := 0
	n := len(nums)
	prefixsum := make([]int, n+1)

	saved := make(map[int]int)
	saved[0] = 1

	for i := 1; i <= n; i++ {
		if nums[i-1]%2 == 1 {
			prefixsum[i] = prefixsum[i-1] + 1
		} else {
			prefixsum[i] = prefixsum[i-1]
		}

		saved[prefixsum[i]]++
		if v, ok := saved[prefixsum[i]-k]; ok {
			ans += v
		}
	}
	return ans
}
