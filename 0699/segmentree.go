package leetcode

import "sort"

type SegmentTree struct {
	dp []int
	n  int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (seg *SegmentTree) Init(n int) {
	seg.dp = make([]int, 4*n)
	seg.n = n
}

func (seg *SegmentTree) Update(l, r, s, t, c, p int) {
	if s == t {
		seg.dp[p] = max(seg.dp[p], c)
		return
	}
	m := s + (t-s)>>1
	if l <= m {
		seg.Update(l, r, s, m, c, 2*p)
	}
	if r > m {
		seg.Update(l, r, m+1, t, c, 2*p+1)
	}
	seg.dp[p] = max(seg.dp[p], max(seg.dp[2*p], seg.dp[2*p+1]))
}

func (seg *SegmentTree) GetMax(l, r, s, t, p int) int {
	if s == t {
		return seg.dp[p]
	}
	m := s + (t-s)>>1
	ans := 0
	if l <= m {
		ans = max(ans, seg.GetMax(l, r, s, m, 2*p))
	}
	if r > m {
		ans = max(ans, seg.GetMax(l, r, m+1, t, 2*p+1))
	}
	return ans
}

func fallingSquares(positions [][]int) []int {
	allnums := make([]int, 0)
	for i := 0; i < len(positions); i++ {
		position := positions[i]
		allnums = append(allnums, position[0], position[0]+position[1]-1)
	}
	// 离散化  allnums, 用 index
	kth := make(map[int]int)
	sort.Ints(allnums)
	k := 1
	kth[allnums[0]] = 1
	for i := 1; i < len(allnums); i++ {
		if allnums[i] != allnums[i-1] { // 去重，且用 index 离散化 allnums
			k++
			kth[allnums[i]] = k
		}
	}
	kl := len(kth)
	seg := &SegmentTree{}
	seg.Init(kl)

	ans := make([]int, 0)

	for i := 0; i < len(positions); i++ {
		start, end, height := kth[positions[i][0]], kth[positions[i][0]+positions[i][1]-1], positions[i][1]
		h := seg.GetMax(start, end, 1, kl, 1)
		seg.Update(start, end, 1, kl, h+height, 1)
		ans = append(ans, seg.GetMax(1, kl, 1, kl, 1))
	}
	return ans
}
