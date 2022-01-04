package leetcode

var cache map[int]int

/* wow,

这道题， 写的超烂啊。

f[3] = 1
f[4] = 2 + 1 = 3
f[5] = 3 + 2 + 1 =  6
f[6] = 4 + 3+ 2 + 1 = 10
.......
 */
 
func numberOfArithmeticSlices(nums []int) int {
	ans := 0
	initialCache(len(nums))
	for i := 0; i < len(nums); i++ {
		if j := i + 1; j < len(nums) {
			dis := nums[i+1] - nums[i]
			k := j + 1
			for ; k < len(nums); k++ {
				if nums[k]-nums[k-1] != dis {
					break
				}
			}
			if k-i >= 3 {
				ans += cache[k-i]
				i = k - 1 - 1
			}
		}
	}

	return ans
}

func initialCache(len int) {
	if len < 4 {
		len = 4
	}

	cache = make(map[int]int, len)
	n := make([]int, len+1)
	n[0], n[1], n[2], n[3] = 0, 0, 0, 1

	for i := 4; i <= len; i++ {
		n[i] = i-2 + n[i-1]
	}

	for i := 3; i <= len; i++ {
		cache[i] = n[i]
	}
}

