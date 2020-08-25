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

	m := make(map[string]string)

	for i := 0; i < len(paths); i++ {
		m[paths[i][0]] = paths[i][1]
	}

	max := 0
	res := ""
	for _, dest := range m {
		d, c := rec(dest, m, 0)
		if c > max {
			max = c
			res = d
		}
	}

	return res
}

func rec(target string, paths map[string]string, count int) (string, int) {
	dest, ok := paths[target]
	if ok {
		return rec(dest, paths, count + 1)
	}
	return target, count
}
