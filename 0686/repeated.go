package leetcode

/*

这道题不是重点， 去看， strStr() 的实现， 了解， KMP 算法
 */
func ceiling(n, d int) int {

	if n != n/d*d {
		return n/d + 1
	}
	return n / d
}

func repeatedStringMatch(a string, b string) int {
	s := ""
	for len(s) < len(a)+len(b)-1 {
		s += a
	}

	for i := 0; i < len(a); i++ {
		if b == s[i:i+len(b)] {
			return ceiling(i+len(b), len(a))
		}
	}
	return -1
}
