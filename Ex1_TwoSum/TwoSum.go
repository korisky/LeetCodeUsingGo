package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

// twoSum using a map to store
func twoSum(nums []int, target int) []int {
	remainMap := make(map[int]int)
	for idx_1 := range nums {
		val_1 := nums[idx_1]
		idx_2, contains := remainMap[val_1]
		if !contains {
			remainMap[target-val_1] = idx_1
		}
		if contains {
			return []int{idx_1, idx_2}
		}
	}
	return []int{}
}
