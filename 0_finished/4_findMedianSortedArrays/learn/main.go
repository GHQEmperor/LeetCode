package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n, m := len(nums1), len(nums2)
	left := (m + n + 1) / 2
	right := (m + n + 2) / 2

	return float64(median(nums1, 0, n-1, nums2, 0, m-1, left)+
		median(nums1, 0, n-1, nums2, 0, m-1, right)) / 2
}

func median(nums1 []int, begin1, end1 int, nums2 []int, begin2, end2, k int) int {
	len1, len2 := end1-begin1+1, end2-begin2+1

	if len1 > len2 {
		return median(nums2, begin2, end2, nums1, begin1, end1, k)
	}
	if len1 == 0 {
		return nums2[begin2+k-1]
	}
	if k == 1 {
		return min(nums1[begin1], nums2[begin2])
	}

	i, j := begin1+min(len1, k/2)-1, begin2+min(len2, k/2)-1
	fmt.Println(i, j)
	if nums1[i] > nums2[j] {
		return median(nums1, begin1, end1, nums2, j+1, end2, k-min(len2, k/2))
	} else {
		return median(nums1, i+1, end1, nums2, begin2, end2, k-min(len1, k/2))
	}

}

func min(v1, v2 int) int {
	if v1 > v2 {
		return v2
	}
	return v1
}
