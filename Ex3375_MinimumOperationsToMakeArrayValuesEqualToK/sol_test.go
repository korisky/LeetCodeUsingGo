package Ex3375_MinimumOperationsToMakeArrayValuesEqualToK

import "testing"

func Test_Func(t *testing.T) {
	arr := make([]int, 0)
	arr = append(arr, 5)
	arr = append(arr, 2)
	arr = append(arr, 5)
	arr = append(arr, 4)
	arr = append(arr, 5)

	println(minOperations(arr, 2))
}

func minOperations(nums []int, k int) int {
	st := make(map[int]bool)
	for _, x := range nums {
		if x < k {
			return -1
		} else if x > k {
			st[x] = true
		}
	}
	return len(st)
}
