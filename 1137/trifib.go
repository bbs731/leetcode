package leetcode

func tribonacci(n int) int {
	a, b, c, d := 0, 1, 1, 0

	if n == 0 {
		return a
	}
	if n == 1 {
		return b
	}
	if n == 2 {
		return c
	}

	for i := 3; i <= n; i ++ {
		d = a + b + c
		a, b, c = b, c, d
	}
	return d
}
