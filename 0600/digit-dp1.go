package leetcode

func findIntegers(n int) int {

	digits := convert2digits(n)
	var getdp func(int, bool, []int, int) int

	/*
	现在面临的困境，是有 dp cache, memory oversize,  没有 dp 会超时！
	这里面的 dp 写的还少了一维， mask  dp[pos][mask] 就会 memory oversize
	因为 mask 最大到了需要覆盖 10^9
	 */

	//dp := make([]int, len(digits))
	//for i := 0; i < len(digits); i++ {
	//	dp[i] = -1
	//}

	getdp = func(pos int, limit bool, digits []int, mask int) int {
		if pos == -1 {
			return 1
		}

		//if !limit && dp[pos] != -1 {
		//	return dp[pos]
		//}

		ans := 0
		var up int
		if limit {
			up = digits[pos]
		} else {
			up = 1
		}

		for i := 0; i <= up; i++ {

			if i == 1 && (mask&(1<<uint(pos+1))) != 0 {
				continue
			}
			ans += getdp(pos-1, limit && i == up, digits, mask|(i<<uint(pos)))

			//if !limit {
			//	dp[pos] = ans
			//}
		}
		return ans
	}

	return getdp(len(digits)-1, true, digits, 0)
}

func convert2digits(n int) []int {

	digits := make([]int, 0)
	for n > 0 {
		if n&1 != 0 {
			digits = append(digits, 1)
		} else {
			digits = append(digits, 0)
		}
		n >>= 1
	}

	return digits
}
