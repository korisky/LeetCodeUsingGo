package main

import (
	"fmt"
	"math"
)

func main() {
	str := "asdbasbuiabweibiuebfleuaksfkjsdbkcjbkjcbkbkzxckSBKDCBKsbduiqwsakdsakdhasjkdhkjasdhjkasdhjksadhjkadhjkasdjkhasjkdhasjkdhjkasdhasjkdhas"
	fmt.Println(lengthOfLongestSubstring(str))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	left, right, max := 0, 0, 0
	for i := 0; i < len(s); i++ {
		for right = left; right < i; right++ {
			if s[right] == s[i] {
				max = int(math.Max(float64(max), float64(i-left)))
				left = right + 1
			}
		}
	}
	return int(math.Max(float64(max), float64(len(s)-left)))
}

func lengthOfLongestSubstring_Map(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	// use char's ascii number as idx, make 1-d array to store
	idxMap := make([]int, 256)
	maxSubLen := 0
	curSubLen := 1
	for i := 0; i < len(s); i++ {
		cha := s[i]
		lastChaPos := idxMap[cha]
		idxMap[cha] = i + 1
		toLastLen := i - lastChaPos + 1
		if toLastLen < curSubLen {
			curSubLen = toLastLen
		}
		if curSubLen > maxSubLen {
			maxSubLen = curSubLen
		}
		curSubLen++
	}
	return maxSubLen
}
