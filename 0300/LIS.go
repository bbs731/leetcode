package leetcode

// This is a O(n^2) solution
// can reduce time complexity to O(nlgn) ?  可以，属于奇淫意巧，看官方的题解吧！
func lengthOfLIS(nums []int) int {
	N := len(nums)
	prev := make([]int, N)   // save the prev element index that's smaller than current
	length := make([]int, N) // save the increasing seq len

	prev[0] = -1
	length[0] = 1

	ans := 1
	for i := 1; i < N; i++ {
		maxlen := 0
		prev[i] = -1
		length[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				if length[j]+1 > maxlen {
					prev[i] = j
					maxlen = length[j] + 1
					length[i] = maxlen
					ans = max(maxlen, ans)
				}
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
