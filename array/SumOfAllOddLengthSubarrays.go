package main

func sumOddLengthSubarrays(arr []int) int {
	sum := 0
	subArrLen := 1
	for subArrLen <= len(arr) {
		for i := 0; i < len(arr); i++ {
			if (i + subArrLen) <= len(arr) {
				for j := i; j < i + subArrLen; j++ {
					sum += arr[j]
				}
			}
		}
		subArrLen += 2
	}
	return sum
}
