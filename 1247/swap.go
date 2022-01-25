package leetcode

func minimumSwap(s1 string, s2 string) int {
	n, m := 0, 0 // 计数  x,y   和 y,x 的数量

	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			continue
		}
		if s1[i] == 'x' {
			n++
		} else {
			m++
		}
	}

	if (n+m)%2 != 0 {
		return -1
	}

	if n%2 == 0 {
		return (n + m) / 2
	}

	return n/2 + m/2 + 2
}
