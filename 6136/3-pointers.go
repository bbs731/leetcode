package leetcode

func arithmeticTriplets(nums []int, diff int) (ans int) {
	i, j := 0, 1
	for _, x := range nums[2:] {
		for nums[j]+diff < x {
			j++
		}
		if nums[j]+diff > x {
			continue
		}
		for nums[i]+diff*2 < x {
			i++
		}
		if nums[i]+diff*2 == x {
			ans++
		}
	}
	return
}
