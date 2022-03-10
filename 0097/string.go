package leetcode

import "fmt"
/*
 	和官网，一样的思路， 为啥，你 写出来，就像是一样！ 
 */

func isInterleave(s1 string, s2 string, s3 string) bool {
	n1 := len(s1)
	n2 := len(s2)

	if len(s3) != n1+n2 {
		return false
	}
	dp := make([][]bool, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]bool, n2+1)
	}
	dp[0][0] = true

	getter := func(i, j int) bool {
		if i < 0 || j < 0 {
			return false
		}
		return dp[i][j]
	}

	gets := func(i int, s string) byte {
		if i < 0 {
			return '$'
		}
		return s[i]
	}

	for i := 0; i <= n1; i++ {
		for j := 0; j <= n2; j++ {
			k := i + j - 1 // s3 current index // 由 s1[i] 和 s2[j] 组成
			if k < 0 {
				continue
			}
			if s3[k] == gets(i-1, s1) {
				dp[i][j] = dp[i][j] || getter(i-1, j)
			}
			if s3[k] == gets(j-1, s2) {
				dp[i][j] = dp[i][j] || getter(i, j-1)
			}
		}
	}

	fmt.Println(dp)
	return dp[n1][n2]

}
