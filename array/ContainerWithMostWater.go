package main

func maxArea(height []int) int {
	res := 0
	for i := 0; i < len(height); i++ {
		h1 := height[i]
		for j := i + 1; j < len(height); j++ {
			h2 := height[j]
			tmp := 0
			if h1 < h2 {
				tmp = h1 * (j - i)
			} else {
				tmp = h2 * (j - i)
			}
			if tmp > res {
				res = tmp
			}
		}
	}
	return res
}

