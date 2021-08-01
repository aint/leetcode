package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
}

var s1Map map[rune]int

func checkInclusion(s1 string, s2 string) bool {
	s1Map = make(map[rune]int, 0)
	for _, r := range s1 {
		s1Map[r] = s1Map[r] + 1
	}

	for i := 0; i <= len(s2) - len(s1); i++ {
		if strings.ContainsRune(s1, rune(s2[i])) {
			if checkPermutations(s1, s2, i) {
				return true
			}
		}
	}
	return false
}

func checkPermutations(s1, s2 string, s2Index int) bool {
	subStrMap := make(map[rune]int, len(s1Map))
	for k, v := range s1Map {
		subStrMap[k] = v
	}

	for i := s2Index; i < s2Index + len(s1); i++ {
		r := rune(s2[i])
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
