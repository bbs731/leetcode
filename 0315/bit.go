package leetcode

import "sort"

/*
same as  offer 51, 逆序对。

离散化，是必须的，不然会超时！
 */
type BIT struct {
	m int
	C []int // index start from 1 to m
}

func lowbit(x int) int {
	return x & (-x)
}

func (b *BIT) add(x int, k int) {
	for x <= b.m {
		b.C[x] += k
		x += lowbit(x)
	}
}

func (b *BIT) getsum(x int) int {
	ans := 0
	for x >= 1 {
		ans += b.C[x]
		x -= lowbit(x)
	}
	return ans
}

func countSmaller(nums []int) []int {
	m := len(nums)
	counts := make([]int, m)
	tmp := make([]int, m) // nums 需要离散化！
	copy(tmp, nums)
	sort.Ints(tmp)

	b := &BIT{
		m,
		make([]int, m+1),
	}

	for i := m - 1; i >= 0; i-- {
		pos := sort.SearchInts(tmp, nums[i])
		b.add(pos+1, 1)
		counts[i] = b.getsum(pos)
	}
	return counts
}
