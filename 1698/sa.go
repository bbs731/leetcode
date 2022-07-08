package leetcode

import "sort"

type SA struct {
	src    string
	sa     []int
	rk     []int
	height []int
}

func NewSA(src string) *SA {
	n := len(src)
	sa := make([]int, n)
	rk := make([]int, 2*n)
	for i := n; i < 2*n; i++ {
		rk[i] = -1
	}

	for i := 0; i < n; i++ {
		rk[i] = int(src[i] - 'a')
		sa[i] = i
	}

	for w := 1; w < n; w <<= 1 {
		sort.Slice(sa, func(x, y int) bool {
			if rk[sa[x]] != rk[sa[y]] {
				return rk[sa[x]] < rk[sa[y]]
			}
			return rk[sa[x]+w] < rk[sa[y]+w]
		})

		oldrk := make([]int, 2*n)
		copy(oldrk, rk)

		rk[sa[0]] = 0
		for i := 1; i < n; i++ {
			if oldrk[sa[i]] == oldrk[sa[i-1]] && oldrk[sa[i]+w] == oldrk[sa[i-1]+w] {
				rk[sa[i]] = rk[sa[i-1]]
			} else {
				rk[sa[i]] = rk[sa[i-1]] + 1
			}
		}
	}

	h := 0
	height := make([]int, n)

	for i := 0; i < n; i++ {
		if rk[i] == 0 {
			continue
		}
		if h > 0 {
			h--
		}
		for j := sa[rk[i]-1]; i+h < n && j+h < n && src[i+h] == src[j+h]; h++ {

		}
		height[rk[i]] = h
	}
	return &SA{src, sa, rk, height}
}

func countDistinct(s string) int {
	sa := NewSA(s)

	n := len(s)
	ans := n * (n + 1) / 2

	for i := 0; i < n; i++ {
		ans -= sa.height[i]
	}
	return ans
}
