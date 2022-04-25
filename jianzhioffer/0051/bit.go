package offer

import "sort"

type BIT struct {
	n int
	C []int
}

func (b *BIT) lowbit(x int) int {
	return x & (-x)
}

func (b *BIT) add(x, v int) {
	for x <= b.n {
		b.C[x] = b.C[x] + v
		x += b.lowbit(x)
	}
}
func (b *BIT) getsum(x int) int {
	ans := 0
	for x >= 1 {
		ans += b.C[x]
		x -= b.lowbit(x)
	}
	return ans
}

/* BIT 最常用来举的一个例子 */

func reversePairs(nums []int) int {

	n := len(nums)
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	sort.Ints(tmp)
	b := BIT{
		n: len(nums),
		C: make([]int, len(nums)+1),
	}
	// 离散化 原 nums 数组， 用他们在数组中的相对位置，作为新值
	for i := 0; i < n; i++ {
		nums[i] = sort.SearchInts(tmp, nums[i]) + 1
	}
	ans := 0
	for i := n - 1; i >= 0; i-- {  // 这个顺序很重要
		ans += b.getsum(nums[i] - 1)
		b.add(nums[i], 1)
	}
	return ans
}
