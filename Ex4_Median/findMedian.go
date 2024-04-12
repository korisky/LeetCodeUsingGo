package main

import "fmt"

func main() {
	//fmt.Println(findMedianSortedArrays([]int{2, 3}, []int{2}))
	//fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{}, []int{2, 3}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	p1, p2, p3 := 0, 0, 0
	mSlice := make([]int, (len(nums1)+len(nums2))/2+1)
	for p1 < len(nums1) && p2 < len(nums2) && p3 < len(mSlice) {
		if nums1[p1] <= nums2[p2] {
			mSlice[p3] = nums1[p1]
			p1++
		} else {
			mSlice[p3] = nums2[p2]
			p2++
		}
		p3++
	}
	for p1 < len(nums1) && p3 < len(mSlice) {
		mSlice[p3] = nums1[p1]
		p1++
		p3++
	}
	for p2 < len(nums2) && p3 < len(mSlice) {
		mSlice[p3] = nums2[p2]
		p2++
		p3++
	}
	fmt.Println(mSlice)

	if (len(nums1)+len(nums2))%2 == 0 {
		if len(mSlice) == 1 {
			return float64(mSlice[0])
		}
		return (float64(mSlice[p3-1]) + float64(mSlice[p3-2])) / 2
	} else {
		return float64(mSlice[p3-1])
	}
}

func findMedianSortedArrays_Improve(nums1 []int, nums2 []int) float64 {

	n, m := len(nums1), len(nums2)
	mSlice := make([]int, (n+m)/2+1)
	even := (n+m)&1 == 0

	p1, p2, p3 := 0, 0, 0
	for p1 < n && p2 < m && p3 < len(mSlice) {
		if nums1[p1] <= nums2[p2] {
			mSlice[p3] = nums1[p1]
			p1++
		} else {
			mSlice[p3] = nums2[p2]
			p2++
		}
		p3++
	}
	for p1 < n && p3 < len(mSlice) {
		mSlice[p3] = nums1[p1]
		p1++
		p3++
	}
	for p2 < m && p3 < len(mSlice) {
		mSlice[p3] = nums2[p2]
		p2++
		p3++
	}

	if even && n+m > 1 {
		return (float64(mSlice[p3-1]) + float64(mSlice[p3-2])) / 2
	}
	return float64(mSlice[p3-1])
}
