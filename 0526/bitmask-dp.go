package leetcode

/*
bitmask dp typical problem!


官网的答案，没咋理解。 自己的挺好理解的。 就是时间差点！
 */
func countArrangement(n int) int {
	size := 1 << uint(n)
	dp := make([]int, size)
	dp[0] = 1
	//ans := 0

	for i := 1; i <= n; i++ {
		for j := 0; j < size; j++ {
			if numofbits(j) != i-1 {
				continue
			}

			if dp[j] == 0 {
				continue
			}

			for pos := 0; pos < n; pos++ {
				if 1<<uint(pos)&j != 0 {
					continue
				}
				if (pos+1)%i == 0 || i%(pos+1) == 0 {
					// number i can put int to pos
					dp[j|1<<uint(pos)] += dp[j]
				}
			}
		}
	}
	return dp[size-1]
}

func numofbits(n int) int {
	cnts := 0
	for n > 0 {
		n &= n - 1
		cnts++
	}
	return cnts
}
