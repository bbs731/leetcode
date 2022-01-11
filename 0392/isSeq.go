package leetcode

func isSubsequence(s string, t string) bool {
	j := 0
	for i := 0; i < len(s); i++ {
		found := false
		for ; j < len(t); {
			if s[i] == t[j] {
				found = true
				j++
				break   // 发现了个有趣的问题， break 的时候，  for statement 里写 j++ 会没有作用
			} else {
				j++
			}
		}
		if !found {
			return false
		}
	}
	return true
}
