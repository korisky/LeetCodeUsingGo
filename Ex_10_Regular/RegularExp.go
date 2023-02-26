package main

import "fmt"

func main() {
	fmt.Println(isMatch("a", "ab*"))
}

func isMatch(s string, p string) bool {
	// init
	mem := make([][]int, len(s))
	for i := range mem {
		mem[i] = make([]int, len(p))
	}
	// dfs
	return dfs(s, p, 0, 0, mem)
}

func dfs(s, p string, sIdx, pIdx int, mem [][]int) bool {
	// if si / pi reach to the end
	if sIdx == len(s) || pIdx == len(p) {
		// p finished, s not
		if sIdx != len(s) {
			return false
		}
		// s finished, p not, following pIdx can be empty (a -> ab*)
		if pIdx != len(p) {
			// only possible way that they could be matched is p following -> a*b*c*d*
			if p[len(p)-1] != '*' {
				return false
			}
			// last char *, also need to check no more regulation
			pIdx++
			for pIdx < len(p) {
				if p[pIdx] != '*' {
					return false
				}
				pIdx += 2
			}
			return true
		}
		// s and p both reached the end simultaneously
		return true
	}

	if mem[sIdx][pIdx] < 0 {
		return false
	}

	// assuming input wildcard will always append to a real char a-z or ./
	// there will be no consecutive *, so we just check pi + 1 for the wildcard
	wild := false
	if pIdx < len(p)-1 && p[pIdx+1] == '*' {
		wild = true
	}

	if wild {
		// if wildcard found, then there are multiple options from 0 matched to n matched
		match := 0
		for {
			// (pIdx, pIdx+1) could match to (sIdx, sIdx + match - 1), then we recursively search from pi+2, si+match
			if dfs(s, p, sIdx+match, pIdx+2, mem) {
				return true
			}
			match++
			// si+match-1 is the point we try match with pi currently
			if sIdx+match-1 >= len(s) {
				break
			}
			// char could not match
			if s[sIdx+match-1] != p[pIdx] && p[pIdx] != '.' {
				break
			}
		}
		mem[sIdx][pIdx] = -1
	}

	if !wild {
		// if no wildcard found, then just need to check current sIdx & pIdx
		if p[pIdx] == '.' || p[pIdx] == s[sIdx] {
			return dfs(s, p, sIdx+1, pIdx+1, mem)
		}
		mem[sIdx][pIdx] = -1
	}

	return false
}
