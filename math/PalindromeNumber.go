package leetcode

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    revertedX := 0

    for tempX := x; tempX != 0; tempX /= 10 {
		revertedX = revertedX * 10 + tempX % 10
    }

    return x == revertedX
}
