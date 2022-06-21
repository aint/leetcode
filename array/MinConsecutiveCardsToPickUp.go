package main

import (
	"math"
)

// 1. Maintain the previous index of each card
// 2. Calc difference between current index and previous (if present) among all pairs
// 3. Return -1 if no duplicate
func minimumCardPickup(cards []int) int {
	min := math.MaxInt32
	m := make(map[int]int)
	for i := 0; i < len(cards); i++ {
		prevIndx, exists := m[cards[i]]
		if exists {
			d := i - prevIndx + 1
			if d < min {
				min = d
			}
		}
		m[cards[i]] = i
	}

	if min == math.MaxInt32 {
		return -1
	}
	return min
}
