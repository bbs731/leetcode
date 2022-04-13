package leetcode

func findIntegers(n int) int {

	// f[pos][0] 代表， pos 位置是 0 结尾的满足条件（无连续 1） 的个数
	f := make([][]int, 32)
	for i := 0; i < 32; i++ {
		f[i] = make([]int, 2)
	}

	f[0][0] = 1
	f[0][1] = 1

	for i := 1; i < 32; i++ {
		f[i][0] = f[i-1][0] + f[i-1][1]
		f[i][1] = f[i-1][0]
	}

	ans := 0
	for i, pre := 31, 0; i >= 0; i-- {
		val := 1 << uint(i)
		if n&val > 0 {
			ans += f[i][0]
			if pre == 1 {
				break
			}
			pre = 1
		} else {
			pre = 0
		}
		// 这个 i == 0  ans + 1 的处理，是真他妈的神奇！ 为什么？
		if i == 0 {
			ans++
		}
	}
	return ans
}
