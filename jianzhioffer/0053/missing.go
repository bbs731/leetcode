package offer

/*
这道题，可以用 二分查找来做， 太巧妙了！
看官方题解吧
 */

func missingNumber(nums []int) int {
	N := len(nums)

	table := make(map[int]struct{}, N)
	for i := 0; i <= N; i++ {
		table[i] = struct{}{}
	}

	for i := 0; i < N; i++ {
		delete(table, nums[i])
	}

	for k := range table {
		return k
	}
	return -1
}
