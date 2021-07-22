package main

import (
	"strings"
)

func multiply(num1 string, num2 string) string {
	// calculate products
	//   2 6
	//   1 6
	//   ---
	// 1 5 6
	// 2 6
	products := make([]string, 0)
	for i := len(num1) - 1; i >= 0; i-- {
		carry := 0
		for j := len(num2) - 1; j >= 0; j-- {
			p := int(num1[i]  - '0') * int(num2[j] - '0') + carry
			if p >= 10 {
				carry = p / 10
				p = p % 10
			} else {
				carry = 0
			}

			index := len(num1) - i
			if (len(products) < index) {
				products = append(products, string(p + '0'))
			} else {
				products[index - 1] = string(p + '0') + products[index - 1]
			}

			if carry > 0 && j == 0 {
				products[index - 1] = string(carry + '0') + products[index - 1]
			}
		}
	}

	// add products
	//   2 6
	//   1 6
	//   ---
	// 1 5 6
	// 2 6 0
	//   ---
	// 4 1 6
	addend1 := products[0]
	for i := 1; i < len(products); i++ {
		// appended or prepended '0' to align addends
		addend2 := products[i] + strings.Repeat("0", i)
		if len(addend2) > len(addend1) {
			addend1 = strings.Repeat("0", len(addend2) - len(addend1)) + addend1
		}
		carry := byte(0)
		for j := len(addend2) - 1; j >= 0; j-- {
			sum := (addend1[j] - '0') + (addend2[j] - '0') + carry
			if sum >= 10 {
				sum = sum % 10
				carry = 1
			} else {
				carry = 0
			}

			addend1 = replaceAtIndex(addend1, sum, j)
			if j == 0 && carry > 0 {
				addend1 = string(carry + '0') + addend1
			}
		}
	}

	// remove leading zeros
	addend1 = strings.TrimLeft(addend1, "0")
	if len(addend1) == 0 {
		return "0"
	}

	return addend1

}

func replaceAtIndex(in string, r byte, i int) string {
	out := []rune(in)
	out[i] = rune(r + '0')
	return string(out)
}
