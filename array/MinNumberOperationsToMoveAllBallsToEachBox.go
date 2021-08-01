package main

func minOperations(boxes string) []int {
	result := make([]int, len(boxes))
	for i := 0; i < len(boxes); i++ {
		ops1 := 0
		curBoxSum1 := 0
		for j := 0; j < i; j++ {
			curBoxSum1 += strBytetoInt(boxes[j])
			ops1 += curBoxSum1
		}
		ops2 := 0
		curBoxSum2 := 0
		for k := len(boxes) - 1; k > i; k-- {
			curBoxSum2 += strBytetoInt(boxes[k])
			ops2 += curBoxSum2
		}
		result[i] = ops1 + ops2
	}
	return result
}

func strBytetoInt(b byte) int {
	if b == '1' {
		return 1
	}
	return 0
}
