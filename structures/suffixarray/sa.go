package main

import (
	"fmt"
	"sort"
)

/*
实现方式，参考：  https://oi-wiki.org/string/sa/
时间复杂度: O(n*lgn*lgn)
 */
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
	oldrk := make([]int, 2*n)

	for i := 0; i < n; i++ {
		rk[i] = int(src[i] - 'a')
		sa[i] = i
	}
	//fmt.Println("sa", sa)

	// binary lifting
	for w := 1; w < n; w = w * 2 {
		sort.Slice(sa, func(x int, y int) bool {
			if rk[sa[x]] != rk[sa[y]] {
				return rk[sa[x]] < rk[sa[y]]
			}
			return rk[sa[x]+w] < rk[sa[y]+w]
		})
		//fmt.Println("sa", sa)
		//fmt.Println("rk", rk)

		// 计算新的 rk 的时候，会有覆盖发生，所以先 copy 一份
		copy(oldrk, rk)
		rk[sa[0]] = 0
		p := 0
		for i := 1; i < n; i++ {
			if oldrk[sa[i]] == oldrk[sa[i-1]] && oldrk[sa[i]+w] == oldrk[sa[i-1]+w] {
				rk[sa[i]] = p
			} else {
				p = p + 1
				rk[sa[i]] = p
			}
		}
	}

	// 根据引理:  height[rk[i]] >= height[rk[i-1]] - 1
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
		for j := sa[rk[i]-1]; src[i+h] == src[j+h]; h++ {
		}
		height[rk[i]] = h
	}

	return &SA{src, sa, rk, height}
}

func main() {

	str := "vamamadn"
	sa := NewSA(str)
	fmt.Println(sa.sa)
	fmt.Println(sa.rk)
	fmt.Println(sa.height)
	for i := 0; i < len(str); i++ {
		fmt.Println(str[sa.sa[i]:])
	}
}
