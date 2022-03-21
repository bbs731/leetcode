package leetcode

type ProductOfNumbers struct {
	nums   []int
	prefix []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		nums:   make([]int, 0),
		prefix: make([]int, 0),
	}
}

func (this *ProductOfNumbers) Add(num int) {
	this.nums = append(this.nums, num)
	n := len(this.nums)

	if num == 0 {
		this.prefix = this.prefix[:0]
	} else {

		if len(this.prefix) == 0 {
			this.prefix = append(this.prefix, num)
		} else {
			l := len(this.prefix)
			this.prefix = append(this.prefix, this.prefix[l-1]*num)
		}
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.prefix)
	if n >= k {
		// we should have prefix[n-k-1] = 1 but n-k-1 index would be -1
		if n == k {
			return this.prefix[n-1]
		}
		return this.prefix[n-1] / this.prefix[n-k-1]
	}
	return 0
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
