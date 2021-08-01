package main

import (
	"strings"
)

var p_map = make(map[rune]int, 0)

func findAnagrams(s string, p string) []int {
	// count chars in `p`
	p_map = make(map[rune]int, 0)
	for _, r := range p {
		if val, ok := p_map[r]; ok {
			p_map[r] = val + 1
		} else {
			p_map[r] = 1
		}
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
	// copy p_map to a new [mutable] one
	substr_map := make(map[rune]int, len(p_map))
	for k, v := range p_map {
		substr_map[k] = v
	}

	for i := sIndex; i < sIndex+pSize && i < len(s); i++ {
		r := rune(s[i])
		if val, ok := substr_map[r]; ok {
			substr_map[r] = val - 1
		} else {
			return false
		}
	}

	for _, v := range substr_map {
		if v != 0 {
			return false
		}
	}

	return true
}
