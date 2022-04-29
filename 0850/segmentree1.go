package leetcode

import (
	"fmt"
	"sort"
)

type Node struct {
	start, end, count, total int
	X                        []int
	Left, Right              *Node
}

func (this *Node) Init(start, end int, nums []int) {
	this.start = start
	this.end = end
	this.X = nums
}
func (this *Node) getLeft() *Node {
	if this.Left == nil {
		m := this.start + (this.end-this.start)>>1
		fmt.Println(this.start, this.end, m) // 这个 GetLeft 有可能会是个无线的循环啊， start = 4, end = 5, m = 4
		this.Left = &Node{this.start, m, 0, 0, this.X, nil, nil}
	}
	return this.Left
}
func (this *Node) getRight() *Node {
	if this.Right == nil {
		m := this.start + (this.end-this.start)>>1
		this.Right = &Node{m, this.end, 0, 0, this.X, nil, nil}
	}
	return this.Right
}

// since we always call update from the root, so the return value would be the total rang query
func (this *Node) Update(l, r, val int) int {
	if l >= r {
		return 0
	}
	if this.start == l && this.end == r {
		this.count += val
	} else {
		m := this.start + (this.end-this.start)>>1
		fmt.Println(this.start, this.end, m)
		this.getLeft().Update(l, min(m, r), val)
		this.getRight().Update(max(m, l), r, val)
	}
	if this.count > 0 {
		this.total = this.X[this.end] - this.X[this.start]
	} else {
		this.total = this.getLeft().total + this.getRight().total
	}
	return this.total
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
	ys := make([]int, 1) // ys 需要多加个 0， 不然， getLeft 会无限循环。
	ys = append(ys, allys[0])
	for i := 1; i < len(allys); i++ {
		if allys[i] != allys[i-1] {
			k++
			kyth[allys[i]] = k
			ys = append(ys, allys[i])
		}
	}

	s := &Node{}
	s.Init(0, len(ys)-1, ys)
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
		cur_y_sum = s.Update(kyth[y1], kyth[y2], typ)
		fmt.Println(cur_y_sum, curx)
		curx = x
	}

	return ans % mod
}
