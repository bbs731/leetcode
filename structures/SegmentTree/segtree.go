package SegmentTree

var n int = 1000

var dp = make([]int, 4*n)
var b = make([]int, 4*n)
var array []int

func build(s, t, p int) {
	if s == t {
		dp[s] = array[s]
	}
	m := s + (t-s)>>1

	build(s, m, 2*p)
	build(m+1, t, 2*p+1)
	dp[p] = dp[2*p] + dp[2*p+1]
}

/* 最基本的版本的 query [l,r] 没有考虑 update 的情况 */
func getsum(l, r int, s, t int, p int) int {
	if l <= s && t <= r {
		return dp[p]
	}
	m := s + (t-s)>>1
	sum := 0
	if l <= m {
		sum += getsum(l, r, s, m, 2*p)
	}
	if r > m {
		sum += getsum(l, r, m+1, t, 2*p+1)
	}
	return sum
}

/*
 [l, r] 区间的数，每个数需要 +c
 */
func update(l, r int, s, t int, c int, p int) {
	if l <= s && t <= r {
		dp[p] += (r - l + 1) * c
		b[p] += c
		return
	}
	m := s + (t-s)>>1

	if b[p] != 0 && s != t { // 这里，为啥要判断  s != t .  其实可以没有
		dp[2*p] += (m - s + 1) * b[p]
		dp[2*p+1] += (t - m) * b[p]
		b[2*p] += b[p]
		b[2*p+1] += b[p]
		b[p] = 0
	}
	if l <= m {
		update(l, r, s, m, c, p*2)
	}
	if r > m {
		update(l, r, m+1, t, c, p*2+1)
	}
	dp[p] = dp[2*p] + dp[2*p+1]
}

func getsum(l, r int, s, t int, p int) int {
	if l <= s && t <= r {
		return dp[p]
	}
	m := s + (t-s)>>1
	if b[p] != 0 {
		dp[p*2] += b[p] * (m - s + 1)
		dp[p*2+1] += b[p] * (t - m)
		b[p*2] += b[p]
		b[p*2+1] += b[p]
		b[p] = 0
	}
	sum := 0
	if l <= m {
		sum += getsum(l, r, s, m, p*2)
	}
	if r > m {
		sum += getsum(l, r, m+1, t, p*2+1)
	}
	return sum
}
