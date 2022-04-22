package leetcode

type NumArray struct {
	d    []int
	b    []int
	nums []int
	n    int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	a := NumArray{
		d:    make([]int, 4*n),
		b:    make([]int, 4*n),
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
		if l <= s && t <= r {
			this.d[p] = (r - l + 1) * c
			this.b[p] = c
			return
		}
		m := s + (t-s)>>1
		if this.b[p] != 0 {
			this.d[p*2] = (m - s + 1) * this.b[p]
			this.d[p*2+1] = (t - m) * this.b[p]
			this.b[p*2] = this.b[p]
			this.b[p*2+1] = this.b[p]
			this.b[p] = 0
		}
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
		if this.b[p] != 0 {
			this.d[p*2] += this.b[p] * (m - s + 1)
			this.d[p*2+1] += this.b[p] * (t - s)
			this.b[p*2] += this.b[p]
			this.b[p*2+1] += this.b[p]
			this.b[p] = 0
		}
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
