package leetcode

import "sort"

type BIT struct {
	n int
	C []int
}

func lowbit(x int) int {
	return x & (-x)
}

func (b *BIT) getsum(x int) int {
	ans := 0
	for x >= 1 {
		ans += b.C[x]
		x -= lowbit(x)
	}
	return ans
}

func (b *BIT) add(x int, k int) {
	for x <= b.n {
		b.C[x] += k
		x += lowbit(x)
	}
}

func reversePairs(nums []int) int {

	n := len(nums)
	double := make([]int, n)
	for i := 0; i < n; i++ {
		double[i] = 2 * nums[i]
	}
	sort.Ints(double)

	b := &BIT{
		n,
		make([]int, n+2),
	}

	ans := 0
	for i := n - 1; i >= 0; i-- {
		// 用数组的 index 来做离散化！
		pos := sort.SearchInts(double, nums[i])
		ans += b.getsum(pos) // pos-1 +1   BIT 的 index 是从 1 开始的
		pos = sort.SearchInts(double, nums[i]*2)
		b.add(pos+1, 1) // BIT  的 index 是从 1 开始的。
	}
	return ans
}
