package main

func flipAndInvertImage(a [][]int) [][]int {
    for _, aa := range a {
		length := len(aa) / 2
		if len(aa) <= 2 {
			length = 1;
		} else if len(aa) % 2 != 0 {
			length += 1
		}
        for i := 0; i < length; i++ {
			tmp := aa[len(aa) - i - 1]
			aa[len(aa) - i - 1] = invert(aa[i])
			aa[i] = invert(tmp)
        }
    }
	return a
}

func invert(v int) int {
	if v == 0 {
		return 1
	}
	return 0
}