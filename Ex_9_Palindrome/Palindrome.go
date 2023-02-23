package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isPalindrome(204))
	fmt.Println(isPalindrome(123))
	fmt.Println(isPalindrome(80))
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(90809))
}

func isPalindrome(x int) bool {
	reverseX := reverseInt(x)
	if reverseX == x {
		return true
	}
	return false
}

func reverseInt(x int) int {
	minInt32 := math.MinInt32
	maxInt32 := math.MaxInt32
	result := 0
	for x != 0 {
		result = result*10 + x%10
		if result < minInt32 || result > maxInt32 {
			return 0
		}
		x /= 10
	}
	return result
}
