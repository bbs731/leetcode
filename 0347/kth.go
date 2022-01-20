package leetcode

type Pair struct {
	num  int
	freq int
}

/*
需要学习一下， 在 golang 中
Heap 的具体用法， 在解决 top K 的问题中，会很有帮助!
 */

func partition(nums []Pair, k int) []Pair {
	pos := 0
	n := len(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i].freq > nums[0].freq {
			continue
		}
		nums[pos+1], nums[i] = nums[i], nums[pos+1]
		pos++
	}
	nums[pos], nums[0] = nums[0], nums[pos]

	if n-pos == k {
		return nums[pos:]
	}

	if n-pos > k {
		return partition(nums[pos+1:], k)
	}
	upper := []Pair{}  // 这里，需要新建个 slice, 直接用 nums[pos:] 有问题。 注意了
	upper = append(upper, nums[pos:]...)
	upper = append(upper, partition(nums[:pos], k-n+pos)...)
	return upper
}

func topKFrequent(nums []int, k int) []int {
	// 先计数， 然后再 partition

	table := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		table[nums[i]]++
	}

	freqs := []Pair{}
	for k, v := range table {
		freqs = append(freqs, Pair{num: k, freq: v})
	}

	top := partition(freqs, k)

	ans := []int{}
	for i := len(top) - 1; i >= 0; i-- {
		ans = append(ans, top[i].num)
	}
	return ans
}
