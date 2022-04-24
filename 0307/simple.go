package leetcode

/*
去掉了，惰性标记的版本， 证明也是可行的. 代码简单了不少！
 */
type NumArray struct {
	d    []int
	nums []int
	n    int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	a := NumArray{
		d:    make([]int, 4*n),
		nums: nums,
	}
	// build segement tree

	var build func(int, int, int)
	build = func(s, t, p int) {
		if s == t {
			a.d[p] = a.nums[s-1] //segement tree index start from 1
			return
		}
		m := s + (t-s)>>1
		build(s, m, p*2)
		build(m+1, t, p*2+1)
		a.d[p] = a.d[p*2] + a.d[p*2+1]
	}

	a.n = n
	build(1, n, 1)
	return a
}

func (this *NumArray) Update(index int, val int) {
	var update func(int, int, int, int, int, int)
	update = func(l, r, s, t, c, p int) {
		if s == t {
			this.d[p] = c
			return
		}
		m := s + (t-s)>>1
		if l <= m {
			update(l, r, s, m, c, p*2)
		}
		if r > m {
			update(l, r, m+1, t, c, p*2+1)
		}
		this.d[p] = this.d[p*2] + this.d[p*2+1]
	}
	update(index+1, index+1, 1, this.n, val, 1)
}

func (this *NumArray) SumRange(left int, right int) int {

	var getsum func(int, int, int, int, int) int
	getsum = func(l, r, s, t, p int) int {
		if l <= s && t <= r {
			return this.d[p]
		}
		m := s + (t-s)>>1
		sum := 0
		if l <= m {
			sum += getsum(l, r, s, m, p*2)
		}
		if r > m {
			sum += getsum(l, r, m+1, t, p*2+1)
		}
		return sum
	}

	return getsum(left+1, right+1, 1, this.n, 1)
}
