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
	rk := make([]int, 2*n) // 开2n 的大小，是为了  sa[i] + w 就不用检查越不越界了
	for i := n; i < 2*n; i++ {
		rk[i] = -1 // 为了，处理，譬如 str = "aaaaaaaa" 的情况
	}
	oldrk := make([]int, 2*n)

	for i := 0; i < n; i++ {
		rk[i] = int(src[i] - 'a')
		sa[i] = i
	}

	// binary lifting
	for w := 1; w < n; w = w * 2 {
		sort.Slice(sa, func(x int, y int) bool {
			if rk[sa[x]] != rk[sa[y]] {
				return rk[sa[x]] < rk[sa[y]]
			}
			return rk[sa[x]+w] < rk[sa[y]+w]
		})

		// 计算新的 rk 的时候，会有覆盖发生，所以先 copy 一份
		copy(oldrk, rk)
		rk[sa[0]] = 0 // rank 的值域应该从 0 开始
		for i := 1; i < n; i++ {
			if oldrk[sa[i]] == oldrk[sa[i-1]] && oldrk[sa[i]+w] == oldrk[sa[i-1]+w] {
				rk[sa[i]] = rk[sa[i-1]]
				//rk[sa[i]] = p
			} else {
				//p++
				//rk[sa[i]] = p
				rk[sa[i]] = rk[sa[i-1]] + 1
			}
		}
	}

	// compute height
	height := make([]int, n)
	h := 0
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

func longestDupSubstring(s string) string {
	sa := NewSA(s)
	result := 0
	idx := -1
	for i := 0; i < len(s); i++ {
		if sa.height[i] > result {
			result = sa.height[i]
			idx = i
		}
	}
	if result == 0 {
		return ""
	}
	return s[sa.sa[idx] : sa.sa[idx]+result]
}
