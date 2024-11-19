package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)
	for i := 0; i < len(s); i++ {

		// Compare the recorded indexes of the current chars from both strings:
		// - If it's the first occurrence of both char, the default index (0) is returned, so 0 == 0.
		// - If both characters have been seen before, their saved indexes must match.
		// - If one character is new and the other is not, so 0 != index.
		r1 := rune(s[i])
		r2 := rune(t[i])
		if m1[r1] != m2[r2] {
			return false
		}

		// Update the index for each char only on its first occurrence since
		// subsequent occurrences don't require updates because the saved indexes remain valid.
		if m1[r1] == 0 {
			m1[r1] = i + 1
		}
		if m2[r2] == 0 {
			m2[r2] = i + 1
		}
	}
	return true
}

func testIsIsomorphic() {
	fmt.Println("true =", isIsomorphic("egg", "add"))
	fmt.Println("true =", isIsomorphic("paper", "title"))
	fmt.Println("false =", isIsomorphic("badc", "baba"))
}
