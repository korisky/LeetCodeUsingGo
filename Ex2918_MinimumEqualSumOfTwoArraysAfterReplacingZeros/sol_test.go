package Ex2918_MinimumEqualSumOfTwoArraysAfterReplacingZeros

import (
	"fmt"
	"testing"
)

func Test_Func(t *testing.T) {
	nums1 := []int{3, 2, 0, 1, 0}
	nums2 := []int{6, 5, 0}
	fmt.Printf("%d\n", minSum(nums1, nums2))
}

func minSum(nums1 []int, nums2 []int) int64 {
	var sum1, sum2 int64
	var zero1, zero2 int

	for _, num := range nums1 {
		sum1 += int64(num)
		if num == 0 {
			sum1++
			zero1++
		}
	}

	for _, num := range nums2 {
		sum2 += int64(num)
		if num == 0 {
			sum2++
			zero2++
		}
	}

	if (zero1 == 0 && sum2 > sum1) || (zero2 == 0 && sum1 > sum2) {
		return -1
	}

	if sum1 > sum2 {
		return sum1
	}
	return sum2
}
