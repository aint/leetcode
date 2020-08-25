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

	maxCount := 0
	result := ""
	for i, pair := range paths {
		dest, count := findDestCityAndSteps(pair[1], paths, 0)
		if count > maxCount {
			maxCount = count
			result = dest
		}

		if len(paths)-i < maxCount {
			break
		}
	}

	return result
}

func findDestCityAndSteps(city string, paths [][]string, count int) (string, int) {
	for _, pair := range paths {
		if pair[0] == city {
			return findDestCityAndSteps(pair[1], paths, count+1)
		}
	}
	return city, count
}
