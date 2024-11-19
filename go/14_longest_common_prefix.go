package main

import (
	"fmt"
	"math"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	// find shortest string
	minStrLen := math.MaxInt
	for _, str := range strs {
		if len(str) < minStrLen {
			minStrLen = len(str)
		}
	}
	// char index for all strings
	var i int
	for i = 0; i < minStrLen; i++ {
		currentChar := strs[0][i]
		// iterate through all strings and compare characters at index 'i'
		for j := 0; j < len(strs); j++ {
			str := strs[j]
			if str[i] != currentChar {
				return strs[0][0:i]
			}
		}
	}

	return strs[0][0:i]
}

func testLongestCommonPrefix() {
	fmt.Println("'a' =", longestCommonPrefix([]string{"ab", "a"}))
	fmt.Println("'fl' =", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println("'' = ", longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
