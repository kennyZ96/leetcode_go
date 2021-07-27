package leetcode_go

//方法一，找第k小的数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	if total%2 == 1 {
		midIndex := total / 2
		return getKthSmallest(nums1, nums2, midIndex+1)
	} else {
		midIndex1 := total/2 - 1
		midIndex2 := total / 2
		return (getKthSmallest(nums1, nums2, midIndex1+1) + (getKthSmallest(nums1, nums2, midIndex2+1))) / 2
	}
}
func getKthSmallest(nums1 []int, nums2 []int, k int) float64 {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return float64(nums2[index2+k-1])
		}
		if index2 == len(nums2) {
			return float64(nums1[index1+k-1])
		}
		if k == 1 {
			return float64(min(nums1[index1], nums2[index2]))
		}

		newIndex1 := min(index1+k/2, len(nums1)) - 1
		newIndex2 := min(index2+k/2, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}
}
