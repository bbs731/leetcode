package leetcode

func isHappy(n int) bool {
	visited := make(map[int]struct{})
	visited[n] = struct{}{}
	for n != 1 {
		sum := 0
		for n != 0 {
			digit := n % 10
			n = n / 10
			sum += digit * digit
		}
		n = sum
		if _, ok := visited[n]; ok {
			return false
		}
		visited[n] = struct{}{}
	}
	return true
}
