package leetcode

import (
	"sort"
)

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

func countRangeSum(nums []int, lower int, upper int) int {
	n := len(nums)
	presum := make([]int, n+1)       // presum[i+1] mean nums[0] + nums[1] + ... + nums[i]
	allnums := make([]int, 1, 3*n+1) // 这里实际上相当于初始化： allnums[0] = 0

	for i := 0; i < n; i++ {
		presum[i+1] = presum[i] + nums[i]
		allnums = append(allnums, presum[i+1])
		allnums = append(allnums, presum[i+1]-upper)
		allnums = append(allnums, presum[i+1]-lower)
	}

	sort.Ints(allnums)

	k := 1
	kth := map[int]int{allnums[0]: k}
	for i := 1; i <= 3*n; i++ {
		if allnums[i] != allnums[i-1] {
			k++
			kth[allnums[i]] = k
		}
	}

	counts := 0
	kl := len(kth)
	b := &BIT{
		make([]int, kl+1),
		make([]int, kl+1),
		kl,
	}
	b.add1(kth[0], kth[0], 1) // 这个初始化，很重要啊！但是是啥意思呢？ 把 0 对应的离散值，设置为 1
	for i := 1; i <= n; i++ {
		l := kth[presum[i]-upper]
		r := kth[presum[i]-lower]
		pos := kth[presum[i]]
		counts += b.getsum1(l, r)
		b.add1(pos, pos, 1)
	}
	return counts
}
