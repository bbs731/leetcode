package leetcode

import "sort"

/*
答案，写的还是太臃肿了!
看看官网，有更简洁的解法！
 */

// With two sorted arrays find the kth element, can do it in log(m+n) time?
func findKth(nums1 []int, nums2 []int, k int) int {
	if len(nums1) == 0 {
		return nums2[k-1]
	}
	if len(nums2) == 0 {
		return nums1[k-1]
	}
	// len() == 1 还需要 check
	if len(nums1) == 1 {
		pos := sort.SearchInts(nums2, nums1[0])
		if k <= pos {
			return nums2[k-1]
		}
		if k == pos+1 {
			return nums1[0]
		}
		return nums2[k-2]
	}

	if len(nums2) == 1 {
		pos := sort.SearchInts(nums1, nums2[0])
		if k <= pos {
			return nums1[k-1]
		}
		if k == pos+1 {
			return nums2[0]
		}
		return nums1[k-2]
	}

	var left, right []int

	if nums1[0] == nums2[0] {
		if k <= 2 {
			return nums1[0]
		}
		k -= 2
		return findKth(nums1[1:], nums2[1:], k)

	} else if nums1[0] < nums2[0] {
		left = nums1
		right = nums2
	} else {
		left = nums2
		right = nums1
	}

	// check whether the two arrays intersect
	if right[0] >= left[len(left)-1] {
		// two array not intersect
		if k <= len(left) {
			return left[k-1]
		} else {
			k -= len(left)
			return right[k-1]
		}
	}

	// now intersect
	rightpos := sort.SearchInts(left, right[0])
	leftpos := sort.SearchInts(right, left[len(left)-1])

	// extream left
	if k <= rightpos {
		return left[k-1]
	}
	if k == rightpos+1 {
		return right[0]
	}

	if k > len(left)+leftpos {
		return right[k-len(left)-leftpos+leftpos-1]
	}

	// else
	k -= rightpos
	return findKth(left[rightpos:], right[:leftpos], k)

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	if (m+n)%2 == 1 {
		k := (m+n)>>1 + 1
		return float64(findKth(nums1, nums2, k))
	}

	l := (m + n) >> 1
	r := (m+n)>>1 + 1
	ln := findKth(nums1, nums2, l)
	rn := findKth(nums1, nums2, r)

	return (float64(ln) + float64(rn)) / float64(2)
}
