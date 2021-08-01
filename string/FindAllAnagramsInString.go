package main

import (
	"strings"
)

var pMap map[rune]int

func findAnagrams(s string, p string) []int {
	// count chars in `p`
	pMap = make(map[rune]int, 0)
	for _, r := range p {
		pMap[r] = pMap[r] + 1
	}

	pSize := len(p)
	result := make([]int, 0)
	for i := 0; i <= len(s) - len(p); i++ {
		if strings.ContainsRune(p, rune(s[i])) {
			if isAnagram(s, i, pSize) {
				result = append(result, i)
			}
		}
	}
	return result
}

func isAnagram(s string, sIndex, pSize int) bool {
	// copy pMap to a new [mutable] one
	subStrMap := make(map[rune]int, len(pMap))
	for k, v := range pMap {
		subStrMap[k] = v
	}

	for i := sIndex; i < sIndex+pSize && i < len(s); i++ {
		r := rune(s[i])
		if val, ok := subStrMap[r]; ok {
			subStrMap[r] = val - 1
		} else {
			return false
		}
	}

	for _, v := range subStrMap {
		if v != 0 {
			return false
		}
	}

	return true
}
