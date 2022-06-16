package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseOnlyLetters("a-bC-dEf-ghIj"))
}

func reverseOnlyLetters(s string) string {
    bs := []byte(s)
    // check if i & j are valid
	i, j := 0, len(bs)-1
	for i < j {
		if !isEnglishLetter(bs[i]) {
			i += 1
			continue
		}

		if !isEnglishLetter(bs[j]) {
			j -= 1
			continue
		}

		t := bs[i]
		bs[i] = bs[j]
		bs[j] = t

		i += 1
		j -= 1
	}

	return string(bs)
}

func isEnglishLetter(c byte) bool {
	return (c >= 65 && c <= 90) || (c >= 97 && c <= 122)
}
