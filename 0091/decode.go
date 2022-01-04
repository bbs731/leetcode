package leetcode

func isValidDigit(a byte) bool {
	if a == '0' {
		return false
	}
	return true
}
func isValid(nums []byte) bool {
	if nums[0] == '1' {
		return true
	} else if nums[0] == '2' {
		return nums[1] >= '0' && nums[1] <= '6'
	}
	return false
}

func numDecodings(s string) int {

	N := len(s)
	dp := make([]int, N+1)
	dp[N] = 1
	ss := []byte(s)

	if N == 1 {
		if isValidDigit(ss[0]) {
			return 1
		}
		return 0
	}

	if isValidDigit(ss[N-1]) {
		dp[N-1] = 1
	}
	/*
		dp[n] = dp[n-1] + dp[n-2]
	 */
	for j := N - 2; j >= 0; j-- {
		if isValidDigit(ss[j]) {
			dp[j] = dp[j+1]
		}
		if isValid(ss[j : j+2]) {
			dp[j] += dp[j+2]
		}
		if dp[j] == 0 {
			// corner case  "2101"
			// look ahead one digit
			if j == 0 || !isValid(ss[j-1:j+1]) {
				break
			}
			// otherwise we continue with j--
		}
	}
	return dp[0]
}
