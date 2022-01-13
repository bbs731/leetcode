package leetcode

/*
有空间复杂度 O(1) 的解法， 想想！  
 */
func merge(nums1 []int, m int, nums2 []int, n int) {

	tmp := make([]int, m+n)
	i := 0
	j := 0
	k := 0

	for ; i < m && j < n; k++ {
		if nums1[i] >= nums2[j] {
			tmp[k] = nums2[j]
			j++
		} else {
			tmp[k] = nums1[i]
			i++
		}
	}

	for i < m {
		tmp[k] = nums1[i]
		k++
		i++
	}

	for j < n {
		tmp[k] = nums2[j]
		j++
		k++
	}

	for i := 0; i < m+n; i++ {
		nums1[i] = tmp[i]
	}
}
