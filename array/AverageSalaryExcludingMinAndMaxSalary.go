package main

func average(salary []int) float64 {
	min, max, sum := salary[0], salary[0], 0
	for _, s := range salary {
		if s < min {
			min = s
		} else if s > max {
			max = s
		}
		sum += s
	}

	sum = sum - (min + max)
	return float64(sum) / float64(len(salary) - 2)
}