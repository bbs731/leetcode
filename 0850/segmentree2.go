package leetcode

import (
	"fmt"
	"sort"
)

/*
 用 数组的这种实现线段树的方法，是不可能写出来的， 参看 Segmentree1.go 的实现吧！
 */
type SegmentTree struct {
	counts []int
	total  []int
	X      []int
	n      int
}

func (seg *SegmentTree) Init(nums []int) {
	seg.X = nums
	n := len(nums)
	seg.counts = make([]int, 4*n)
	seg.total = make([]int, 4*n)
	seg.n = n
}

func (seg *SegmentTree) Update(l, r, s, t, c, p int) {
	if l <= s && t <= r {
		seg.counts[p] += c
		return
	}
	//if s == t {
	//	seg.counts[p] += c
	//	return
	//}
	m := s + (t-s)>>1
	if l <= m {
		seg.Update(l, r, s, m, c, 2*p)
	}
	if r > m {
		seg.Update(l, r, m+1, t, c, 2*p+1)
	}
	seg.total[p] = seg.total[2*p] + seg.total[2*p+1]
}

func (seg *SegmentTree) Query(l, r, s, t, p int) int {
	if l <= s && t <= r {
		if seg.counts[p] > 0 {
			return seg.X[t-1] - seg.X[s-1]
		}
		return 0
	}

	value := 0
	m := s + (t-s)>>1
	if l <= m {
		value += seg.Query(l, r, s, m, 2*p)
	}
	if r > m {
		value += seg.Query(l, r, m+1, t, 2*p+1)
	}
	return value
}

func rectangleArea(rectangles [][]int) int {

	mod := 1000000007
	OPEN := 1
	CLOSE := -1
	events := make([][]int, 0, len(rectangles)*2)
	allys := make([]int, 0)

	for _, rectangle := range rectangles {
		events = append(events, []int{rectangle[0], OPEN, rectangle[1], rectangle[3]})
		events = append(events, []int{rectangle[2], CLOSE, rectangle[1], rectangle[3]})
		allys = append(allys, rectangle[1])
		allys = append(allys, rectangle[3])
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	sort.Ints(allys)
	kyth := make(map[int]int)
	k := 1
	kyth[allys[0]] = 1
	ys := make([]int, 0) // ys 需要多加个 0， 不然， getLeft 会无限循环。
	ys = append(ys, allys[0])
	for i := 1; i < len(allys); i++ {
		if allys[i] != allys[i-1] {
			k++
			kyth[allys[i]] = k
			ys = append(ys, allys[i])
		}
	}

	s := &SegmentTree{}
	s.Init(ys)
	fmt.Println(ys)

	ans := 0
	curx := events[0][0]
	cur_y_sum := 0

	for _, event := range events {
		x := event[0]
		typ := event[1]
		y1 := event[2]
		y2 := event[3]
		ans += cur_y_sum * (x - curx)
		s.Update(kyth[y1], kyth[y2], 1, len(ys), typ, 1)
		//cur_y_sum = s.Update(kyth[y1], kyth[y2], typ)
		cur_y_sum = s.Query(1, len(ys), 1, len(ys), 1)
		fmt.Println(cur_y_sum)
		curx = x
	}

	return ans % mod
}
