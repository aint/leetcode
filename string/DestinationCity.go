package main

import "fmt"

func main() {

	input1 := [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}

	fmt.Println("RESULT1 = ", destCity(input1))

	input2 := [][]string{{"A", "Z"}}

	fmt.Println("RESULT2 = ", destCity(input2))
}

func destCity(paths [][]string) string {
	if len(paths) == 1 {
		return paths[0][1]
	}

	max := 0
	res := ""
	for i, pair := range paths {
		d, count := rec(pair[1], paths, 0)
		if count > max {
			max = count
			res = d
		}

		if len(paths)-i < max {
			break
		}
	}

	return res
}

func rec(target string, paths [][]string, count int) (string, int) {
	for _, pair := range paths {
		if pair[0] == target {
			return rec(pair[1], paths, count+1)
		}
	}
	return target, count
}
