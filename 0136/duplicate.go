package leetcode

func singleNumber(nums []int) int {

	hash := make(map[int]struct{})

	for i := 0; i < len(nums); i++ {
		if _, ok := hash[nums[i]]; ok {
			delete(hash, nums[i])
		} else {
			hash[nums[i]] = struct{}{}
		}
	}

	for k := range hash {
		return k
	}
	return -1
}
