package main

func isAnagram(s string, t string) bool {
    m := make(map[rune]int, 0)
    for _, c := range s {
        if val, ok := m[c]; ok {
            m[c] = val + 1
        } else {
            m[c] = 1
        }
    }
    for _, c := range t {
        if val, ok := m[c]; ok {
            m[c] = val - 1;
        } else {
            return false
        }
    }

    for _, v := range m {
        if v > 0 {
            return false
        }
    }

    return true

}