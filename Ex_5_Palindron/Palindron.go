package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("abaaa"))
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	// init for dynamic programming
	table := make([][]int, len(s))
	for i := range table {
		table[i] = make([]int, len(s))
	}
	maxLen, maxStartIdx := 1, 0
	for i := 0; i < len(table); i++ {
		table[i][i] = 1 // single character is a len 1 valid palindrome
		maxStartIdx = i
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			table[i][i+1] = 1 // e.g. "bb" is also palindrome
			maxLen = 2
			maxStartIdx = i
		}
	}
	curLenP := 3
	for curLenP <= len(s) {
		for i := 0; i < len(s)+1-curLenP; i++ {
			j := i + curLenP - 1
			if s[i] == s[j] && table[i+1][j-1] == 1 {
				table[i][j] = 1
				if maxLen < curLenP {
					maxLen = curLenP
					maxStartIdx = i
				}
			}
		}
		curLenP++
	}
	return s[maxStartIdx : maxStartIdx+maxLen]
}
