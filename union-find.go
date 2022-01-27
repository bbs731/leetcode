package leetcode

func main() {

	/*
			union-find template
	 */
	var n int
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int

	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(from, to int) {
		x, y := find(from), find(to)
		parent[x] = y
	}

}
