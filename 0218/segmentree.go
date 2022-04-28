package leetcode

import "sort"

type SegmentTree struct {
	n  int
	dp []int
}

func (seg *SegmentTree) Init(nums []int) {
	n := len(nums)
	seg.n = n
	seg.dp = make([]int, 4*n)

	var build func(int, int, int)
	build = func(s, t, p int) {
		if s == t {
			seg.dp[s] = nums[s-1] // nums index start from 0 and dp index start from 1
			return
		}
		m := s + (t-s)>>1
		build(s, m, 2*p)
		build(m+1, t, 2*p+1)
		seg.dp[p] = max(seg.dp[p*2+1], seg.dp[2*p])
	}
}

func (seg *SegmentTree) update(l, r, s, t, c, p int) {
	if s == t {
		seg.dp[p] = max(seg.dp[p], c)
		return
	}
	m := s + (t-s)>>1
	if l <= m {
		seg.update(l, r, s, m, c, 2*p)
	}
	if r > m {
		seg.update(l, r, m+1, t, c, 2*p+1)
	}
	seg.dp[p] = max(seg.dp[p], max(seg.dp[2*p+1], seg.dp[2*p]))
}

/*
range max query [l, r]
 */
func (seg *SegmentTree) GetMax(l, r, s, t, p int) int {
	if s == t {
		return seg.dp[p]
	}
	value := 0
	m := s + (t-s)>>1
	if l <= m {
		value = max(value, seg.GetMax(l, r, s, m, 2*p))
	}
	if r > m {
		value = max(value, seg.GetMax(l, r, m+1, t, 2*p+1))
	}
	return value
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getSkyline(buildings [][]int) [][]int {
	n := len(buildings)
	allxs := make([]int, 0)
	lines := make(map[int]bool)

	for i := 0; i < n; i++ {
		allxs = append(allxs, buildings[i][0]*10, buildings[i][0]*10+1, buildings[i][1]*10, buildings[i][1]*10+1) // 这是个技巧，
		// 用来，区别这种情况 ： [[4,9,10],[4,9,15],[4,9,12],[10,12,10],[10,12,8]]  譬如 9, 10 之间加个间隔。
		// [[2,4,70],[3,8,30],[6,100,41],[7,15,70],[10,30,102],[15,25,76],[60,80,91],[70,90,72],[85,120,59]]  6，和 7 之间，需要有个数来隔开
		lines[buildings[i][0]*10] = true
		lines[buildings[i][1]*10] = true
	}
	sort.Ints(allxs)
	//nums := make([]int, 0)
	kth := make(map[int]int)
	k := 1
	kth[allxs[0]] = 1
	//nums = append(nums, allxs[0])
	// 离散化 allxs
	for i := 1; i < len(allxs); i++ {
		if allxs[i] != allxs[i-1] {
			k++
			kth[allxs[i]] = k
		}
	}
	kl := len(kth)
	array := make([]int, kl)
	seg := &SegmentTree{}
	seg.Init(array)

	for i := 0; i < n; i++ {
		l := kth[buildings[i][0]*10]
		r := kth[buildings[i][1]*10]
		seg.update(l, r, 1, kl, buildings[i][2], 1)
	}

	ans := make([][]int, 0)

	prev := allxs[0]
	previ := kth[allxs[0]]
	prevh := seg.GetMax(previ, previ, 1, kl, 1)
	ans = append(ans, []int{prev / 10, prevh})

	for i := 1; i < len(allxs); i++ {
		cur := allxs[i]
		if lines[cur] {
			curi := kth[cur]
			curh := seg.GetMax(curi+1, curi+1, 1, kl, 1)
			prevh = seg.GetMax(curi-1, curi-1, 1, kl, 1)
			if curh != prevh {
				if cur == ans[len(ans)-1][0]*10 && curh == ans[len(ans)-1][1] {
					continue
				}
				ans = append(ans, []int{cur / 10, curh})
			}
		}
	}
	return ans
}
