package main

import "fmt"

// segmentation tree 的 index 取值范围是: [1，n]
// query 的时候也是 l, r 的取值范围是闭区间 [1, n],  index 不是从 0 开始的， 因此 mid = (l+r)/2 下取整

type seg []int // save sum, 区间和

func (t seg) add(o, l, r, idx, val int) {
	if l == r {
		t[o] += val
		return
	}
	mid := (l + r) >> 1

	if idx <= mid {
		t.add(o*2, l, mid, idx, val)
	} else {
		t.add(o*2+1, mid+1, r, idx, val)
	}
	t[o] = t[o*2] + t[o*2+1]
	return
}

func (t seg) query(o, l, r, L, R int) int {
	if L <= l && r <= R {
		return t[o]
	}

	sum := 0
	mid := (l + r) / 2
	if L <= mid {
		sum += t.query(2*o, l, mid, L, R)
	}
	if R > mid {
		sum += t.query(2*o+1, mid+1, r, L, R)
	}
	return sum

}

func (t *seg) NewSegmentTree(n int) {
	*t = make([]int, 4*n)
}

func main() {
	var t seg
	t.NewSegmentTree(5) // 初始化线段树，数组大小为 5, index 的范围是: [1,5]

	// tree root 的 index (数组的下标）为 1,   （跟 heap 用数组表示的时候类似)
	t.add(1, 1, 5, 1, 1) // 在索引 0 处添加值 1
	t.add(1, 1, 5, 3, 3) // 在索引 2 处添加值 3

	fmt.Println(t.query(1, 1, 5, 1, 3)) // 输出 4，查询索引 0 到 2 的区间和
	fmt.Println(t.query(1, 1, 5, 1, 3)) // 输出 4，查询索引 1 到 3 的区间和
	fmt.Println(t.query(1, 1, 5, 1, 4)) // 输出 4，查询索引 0 到 4 的区间和
	fmt.Println(t.query(1, 1, 5, 4, 5)) // 输出 0，查询索引 3 到 4 的区间和（没有添加值）
	fmt.Println(t.query(1, 1, 5, 3, 3)) // 输出 3，查询索引 2 的值
}
