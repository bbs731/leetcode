package leetcode

import (
	"fmt"
	"sort"
)

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
		rk[sa[0]] = 0 // rank 的值域，应该从 0 开始 , 或者 p = 0 开始
		for i := 1; i < n; i++ {
			if oldrk[sa[i]] == oldrk[sa[i-1]] && oldrk[sa[i]+w] == oldrk[sa[i-1]+w] {
				rk[sa[i]] = rk[sa[i-1]] // 不能区分，所以，值域不增加
				//rk[sa[i]] = p
			} else {
				//p++
				//rk[sa[i]] = p
				rk[sa[i]] = rk[sa[i-1]] + 1
			}
		}

		//fmt.Println("sa", sa)
		//fmt.Println("rk", rk)
	}

	// 根据引理:  height[rk[i]] >= height[rk[i-1]] - 1.   证明在这里： https://oi-wiki.org/string/sa/
	/*
		推倒的过程，
		根据定义：  height[i] = lcp( suffix(sa[i]), suffix(sa[i-1])
		所以：
			height[rk[i]] = lcp (suffix(sa[rk[i]]),  suffix(sa[rk[i]-1))
	                      = lcp (suffix(i), suffix(sa[rk[i]-1]))

		令， height[rk[i-1]] = h,   则  height[rk[i]] >= h -1
	    令， j = sa[rk[i]-1]  则，

	     height[rk[i]] = lcp(suffix(i), suffix(j))
	                   = h-1 + lcp(suffix(i+h-1），suffix(j+h-1))
		为以下代码，计算 height[rk[i]] 的步骤过程
	 */
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

func longestRepeatingSubstring(s string) int {
	sa := NewSA(s)

	result := 0
	for i := 0; i < len(s); i++ {
		result = max(result, sa.height[i])
	}
	fmt.Println(sa.height)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
