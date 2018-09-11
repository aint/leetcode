package main

import (
    "fmt"
)

func main() {
    fmt.Println(isPalindrome(2147483647))
}

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    var revertedX int
    tempX := x

    for {
        remainder := tempX % 10
        revertedX = revertedX * 10 + remainder

        tempX = tempX / 10
        if tempX == 0 {
            break
        }
    }

    fmt.Println("revertedX", revertedX)
    fmt.Println("x", x)
    fmt.Println("tempX", tempX)

    return x == revertedX
}
