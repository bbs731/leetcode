package leetcode

/*
无论是用 BIT还是 segmentree 的实现，都会超时，

因为每次都需要重新建树。 可知，线段树，还是比较适合，用来做离线的数据处理。 即： 当，所有的 book.start, book.end 的值已知的情况下。

对于在线的，这种问题， sorted list 的实现， 也就是 Java 中的 tree map 可能是个更好的选择

 */

import "sort"

type BIT struct {
	t1 []int // 保存的是 差分数组 bi  // index 是 从  1 到 n
	t2 []int // 保存的是  i*bi 差分数组
	n  int
}

func lowbit(x int) int {
	return x & (-x)
}

func (b *BIT) add(x, v int) {
	v1 := x * v
	for x <= b.n {
		b.t1[x] += v
		b.t2[x] += v1
		x += lowbit(x)
	}
}

func getsum(t []int, x int) int {
	ret := 0
	for x >= 1 {
		ret += t[x]
		x -= lowbit(x)
	}
	return ret
}

/* 注意， 调用的时候，需要使用 add1, getsum1,  不能直接使用 add, 和 getsum
 */

func (b BIT) add1(l, r int, v int) {
	b.add(l, v)
	b.add(r+1, -v)
}

func (b BIT) getsum1(l, r int) int {

	// sum(i ~r) bi *(r+1) -  sum(i ~r)bi*i
	return (r+1)*getsum(b.t1, r) - (l-1+1)*getsum(b.t1, l-1) - (getsum(b.t2, r) - getsum(b.t2, l-1))
}

type MyCalendar struct {
	queries [][2]int
}

func Constructor() MyCalendar {

	return MyCalendar{
		queries: make([][2]int, 0),
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
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
	b := &BIT{
		make([]int, kl+1),
		make([]int, kl+1),
		kl,
	}

	for i := 0; i < n-1; i++ {
		l := kth[this.queries[i][0]]
		r := kth[this.queries[i][1]-1]
		b.add1(l, r, 1)
	}

	if b.getsum1(kth[this.queries[n-1][0]], kth[this.queries[n-1][1]-1]) != 0 {
		// remove the last element
		this.queries = this.queries[:n-1]
		return false
	}
	return true
}
