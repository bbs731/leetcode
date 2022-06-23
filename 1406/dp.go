package leetcode

/*
走狗屎运了吗？ 一次就过！
 */
func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)

	alice := make([]int, n+1)
	bob := make([]int, n+1)

	// 初始化
	bob[n-1] = 0
	alice[n-1] = stoneValue[n-1]

	for i := n - 2; i >= 0; i-- {
		// take one, two , three
		alice[i] = bob[i+1] + stoneValue[i]
		bob[i] = alice[i+1]

		if i+2 <= n {
			a := stoneValue[i] + stoneValue[i+1] + bob[i+2]
			if a > alice[i] {
				alice[i] = a
				bob[i] = alice[i+2]
			}
		}

		if i+3 <= n {
			a := stoneValue[i] + stoneValue[i+1] + stoneValue[i+2] + bob[i+3]
			if a > alice[i] {
				alice[i] = a
				bob[i] = alice[i+3]
			}
		}
	}

	if alice[0] > bob[0] {
		return "Alice"
	} else if alice[0] < bob[0] {
		return "Bob"
	}
	return "Tie"

}
