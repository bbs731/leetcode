package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func profit(prices []int) int {
	ans := 0
	N := len(prices)
	min := prices[0]

	for i := 1; i < N; i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			ans = max(ans, prices[i]-min)
		}
	}
	return ans
}

type position struct {
	start int
	end   int
}

var cache map[position]int

func recursive(prices []int, start, end int) int {

	res := 0
	if val, ok := cache[position{start, end}]; ok {
		return val
	}

	res = profit(prices[start : end+1])

	// in case length is great than 3, we can split
	for k := start + 1; k < end-1; k++ {
		m1 := recursive(prices, start, k)
		m2 := recursive(prices, k+1, end)
		res = max(m1+m2, res)
	}

	cache[position{start, end}] = res
	return res
}

func maxProfit_timeout(prices []int) int {
	N := len(prices)
	cache = make(map[position]int)
	return recursive(prices, 0, N-1)
}
