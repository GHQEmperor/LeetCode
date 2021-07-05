package findmediansortedarrays

/*
	每次删除 k/2		 k = k - min(k/2,oneLen).
	从每个数组中同时抽出第 k/2 个，比较大小，谁小删谁.
*/

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	n, m := len(nums1), len(nums2)
	left := (n + m + 1) / 2
	right := (n + m + 2) / 2
	return float64(getKth(nums1, nums2, 0, n-1, 0, m-1, left)+
		getKth(nums1, nums2, 0, n-1, 0, m-1, right)) / 2
}

func getKth(nums1, nums2 []int, start1, end1, start2, end2, k int) int {
	len1, len2 := end1-start1+1, end2-start2+1

	if len1 > len2 {
		return getKth(nums2, nums1, start2, end2, start1, end1, k)
	}
	if len1 == 0 {
		return nums2[start2+k-1]
	}
	if k == 1 {
		return min(nums1[start1], nums2[start2])
	}

	i := start1 + min(len1, k/2) - 1
	j := start2 + min(len2, k/2) - 1

	if nums1[i] > nums2[j] {
		return getKth(nums1, nums2, start1, end1, j+1, end2, k-min(len2, k/2))
	}
	return getKth(nums1, nums2, i+1, end1, start2, end2, k-min(len1, k/2))

}

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}
