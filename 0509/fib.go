package leetcode

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b, c := 0, 1, 0
	for i := 2; i <= n; i++ {
		c = a + b
		a, b = b, c
	}
	return c
}
