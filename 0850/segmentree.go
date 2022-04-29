package leetcode

import (
	"fmt"
	"sort"
)

type SegmentTree struct {
	dp []int
	n  int
}

func (seg *SegmentTree) Init(n int) {
	seg.dp = make([]int, 4*n)
	seg.n = n
}

func (seg *SegmentTree) Update(l, r, s, t, c, p int) {
	if s == t {
		seg.dp[p] = c
		return
	}
	m := s + (t-s)>>1
	if l <= m {
		seg.Update(l, r, s, m, c, 2*p)
	}
	if r > m {
		seg.Update(l, r, m+1, t, c, 2*p+1)
	}
	seg.dp[p] = seg.dp[2*p] + seg.dp[2*p+1]
}

func (seg *SegmentTree) GetSum(l, r, s, t, p int) int {
	if s == t {
		return seg.dp[p]
	}
	value := 0
	m := s + (t-s)>>1
	if l <= m {
		value += seg.GetSum(l, r, s, m, 2*p)
	}
	if r > m {
		value += seg.GetSum(l, r, m+1, t, 2*p+1)
	}
	return value
}

func rectangleArea(rectangles [][]int) int {

	mod := 1000000007
	allxs := make([]int, 0)
	allys := make([]int, 0)
	realys := make([]int, 0)
	valid := make(map[int]bool)

	for i := 0; i < len(rectangles); i++ {
		allxs = append(allxs, rectangles[i][0], rectangles[i][2])
		allys = append(allys, rectangles[i][1]*10, rectangles[i][1]*10+1, rectangles[i][3]*10, rectangles[i][3]*10+1)
		valid[rectangles[i][1]*10] = true
		valid[rectangles[i][3]*10] = true
	}

	sort.Ints(allxs)
	sort.Ints(allys)
	// 离散化， 横坐标和纵坐标
	kxth := make(map[int]int)
	kyth := make(map[int]int)
	rkxth := make(map[int]int)
	rkyth := make(map[int]int)

	k := 1
	kxth[allxs[0]] = 1
	rkxth[1] = allxs[0]
	for i := 1; i < len(allxs); i++ {
		if allxs[i] != allxs[i-1] {
			k++
			kxth[allxs[i]] = k
			rkxth[k] = allxs[i]
		}
	}
	kyth[allys[0]] = 1
	rkyth[1] = allys[0]
	k = 1
	realys = append(realys, 1)
	for i := 1; i < len(allys); i++ {
		if allys[i] != allys[i-1] {
			k++
			kyth[allys[i]] = k
			rkyth[k] = allys[i]
			if valid[allys[i]] {
				realys = append(realys, k)
			}
		}
	}

	kyl := len(kyth)
	kxl := len(kxth) //starting from index  1
	// then we create  kxl -1 个 segment tree, index  [2, kxl] 来记录，y 轴上，每个 rectangle 贡献的面积
	segmentrees := make([]*SegmentTree, kxl+1)
	for i := 0; i <= kxl; i++ {
		segmentrees[i] = &SegmentTree{}
		segmentrees[i].Init(kyl)
	}

	for i := 0; i < len(rectangles); i++ {
		xs, xe := kxth[rectangles[i][0]], kxth[rectangles[i][2]]
		ys, ye := kyth[rectangles[i][1]*10], kyth[rectangles[i][3]*10]

		for j := xs; j <= xe-1; j++ {
			fmt.Printf("update: %d, %d, %d\n", j, ys, ye)
			segmentrees[j].Update(ys, ye, 1, kyl, 1, 1)
		}
	}

	ans := 0
	fmt.Println(allys)
	fmt.Println(realys)
	fmt.Printf("rkyth: %v\n", rkyth)
	fmt.Printf("rkxth: %v\n", rkxth)

	/*
	这里涉及到一个问题，就是， 离散化了， x, y 轴上的信息， 但是最后，求面积的时候， 又需要 restore scale 这个太麻烦了，
	也就是说， y 轴上相邻的 index, 我们存的是 1 值，但是算面积的时候， 用用 rkyth[j] - rkyth[j-1] 连计算真实的长度。
	 */

	/*
	这里用到了线段树，但是是个摆设，最终的复杂度是 O(n^2*lgn) ， 看看官网上 O(nlgn) 的解法吧， 没看懂。
	 */

	for i := 1; i <= kxl-1; i++ {
		wide := rkxth[i+1] - rkxth[i]
		for j := 1; j < len(realys); j++ {
			height := (rkyth[realys[j]] - rkyth[realys[j-1]]) / 10
			unit := segmentrees[i].GetSum(realys[j-1], realys[j], 1, kyl, 1)
			if unit > 2 { //  这种 trick 太多了， 属于程序的不定， 应该不是正确的解法！
				ans += height * wide
				ans %= mod
				fmt.Println(i, unit, realys[j-1], realys[j], height, wide)
			}

		}
	}
	return ans
}
