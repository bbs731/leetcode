package leetcode

/*
作者：灵茶山艾府
链接：https://leetcode.cn/problems/find-building-where-alice-and-bob-can-meet/solutions/2533058/chi-xian-zui-xiao-dui-pythonjavacgo-by-e-9ewj/
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

视屏的讲解： https://t.bilibili.com/865609585771151362?spm_id_from=444.41.0.0

线段树的介绍： 视频  https://www.bilibili.com/video/BV18t4y1p736/?spm_id_from=333.788&vd_source=84c3c489cf545fafdbeb3b3a6cd6a112

*/

type seg []int

func (t seg) build(a []int, o, l, r int) {
	if l == r {
		t[o] = a[l-1]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t[o] = max(t[o<<1], t[o<<1|1])
}

// 返回 [L,n] 中 > v 的最小下标（前三个参数表示线段树的节点信息）
func (t seg) query(o, l, r, L, v int) int {
	if v >= t[o] { // 最大值 <= v，没法找到 > v 的数
		return 0
	}
	if l == r { // 找到了
		return l
	}
	m := (l + r) >> 1
	if L <= m {
		pos := t.query(o<<1, l, m, L, v) // 递归左子树
		if pos > 0 { // 找到了
			return pos
		}
	}
	return t.query(o<<1|1, m+1, r, L, v) // 递归右子树
}

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	n := len(heights)
	t := make(seg, n*4)
	t.build(heights, 1, 1, n)
	ans := make([]int, len(queries))
	for qi, q := range queries {
		i, j := q[0], q[1]
		if i > j {
			i, j = j, i
		}
		if i == j || heights[i] < heights[j] {
			ans[qi] = j
		} else {
			pos := t.query(1, 1, n, j+1, heights[i])
			ans[qi] = pos - 1 // 不存在时刚好得到 -1
		}
	}
	return ans
}
