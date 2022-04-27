package leetcode

import "sort"

type segmentree struct {
	n     int
	array []int
	dp    []int
	b     []int
}

func (seg *segmentree) init(nums []int) {

	seg.array = nums
	n := len(seg.array)
	seg.n = n
	seg.dp = make([]int, 4*n)
	seg.b = make([]int, 4*n)
}

func (seg *segmentree) build(s, t, p int) {
	if s == t {
		seg.dp[s] = seg.array[s-1]
		return
	}
	m := s + (t-s)>>1
	seg.build(s, m, 2*p)
	seg.build(m+1, t, 2*p+1)
	seg.dp[p] = seg.dp[2*p] + seg.dp[2*p+1]
}

//func (seg *segmentree) getsum(l, r int, s, t int, p int) int {
//	if l <= s && t <= r {
//		return seg.dp[p]
//	}
//	m := s + (t-s)>>1
//	if seg.b[p] != 0 {
//		seg.dp[p*2] += seg.b[p] * (m - s + 1)
//		seg.dp[p*2+1] += seg.b[p] * (t - m)
//		seg.b[p*2] += seg.b[p]
//		seg.b[p*2+1] += seg.b[p]
//		seg.b[p] = 0
//	}
//	sum := 0
//	if l <= m {
//		sum += seg.getsum(l, r, s, m, 2*p)
//	}
//	if r > m {
//		sum += seg.getsum(l, r, m+1, t, 2*p+1)
//	}
//	return sum
//}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (seg *segmentree) getmax(l, r, s, t, p int) int {
	if s == t {
		return seg.dp[p]
	}
	m := s + (t-s)>>1
	if seg.b[p] != 0 {
		seg.dp[p*2] += seg.b[p] * (m - s + 1)
		seg.dp[p*2+1] += seg.b[p] * (t - m)
		seg.b[p*2] += seg.b[p]
		seg.b[p*2+1] += seg.b[p]
		seg.b[p] = 0
	}
	value := 0
	if l <= m {
		value = max(value, seg.getmax(l, r, s, m, 2*p))
	}
	if r > m {
		value = max(value, seg.getmax(l, r, m+1, t, 2*p+1))
	}
	return value
}
func (seg *segmentree) update(l, r, s, t, c, p int) {
	if s == t {
		seg.dp[p] += c
		return
	}
	m := s + (t-s)>>1

	if l <= m {
		seg.update(l, r, s, m, c, p*2)
	}
	if r > m {
		seg.update(l, r, m+1, t, c, p*2+1)
	}
	seg.dp[p] = seg.dp[p*2] + seg.dp[p*2+1]
}

type MyCalendarThree struct {
	queries [][2]int
}

func Constructor() MyCalendarThree {

	return MyCalendarThree{
		queries: make([][2]int, 0),
	}
}

func (this *MyCalendarThree) Book(start int, end int) int {

	this.queries = append(this.queries, [2]int{start, end})
	// build up segment tree
	n := len(this.queries)

	allnums := make([]int, 0, 2*n)
	for i := 0; i < n; i++ {
		allnums = append(allnums, this.queries[i][0], this.queries[i][1]-1)
	}
	sort.Ints(allnums)

	kth := make(map[int]int)

	k := 1
	kth[allnums[0]] = k
	for i := 1; i < 2*n; i++ {
		if allnums[i] != allnums[i-1] {
			k++
			kth[allnums[i]] = k
		}
	}

	kl := len(kth)
	seg := &segmentree{
	}
	array := make([]int, kl)
	seg.init(array)
	seg.build(1, len(array), 1)

	for i := 0; i < n; i++ {
		l := kth[this.queries[i][0]]
		r := kth[this.queries[i][1]-1]
		seg.update(l, r, 1, kl, 1, 1)
	}

	return seg.getmax(1, kl, 1, kl, 1)

}
