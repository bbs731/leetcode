package leetcode

import "fmt"

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	a := NumArray{}
	sums := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Println(i)
		sums[i] = sums[i-1] + nums[i-1]
	}
	a.sums = sums
	return a
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sums[right+1] - this.sums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
