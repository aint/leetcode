fun isPalindrome(s: String): Boolean {
    var l = 0
    var r = s.length - 1
    while (l <= r) {
        val c1 = s[l].lowercaseChar()
        val c2 = s[r].lowercaseChar()
        if (!c1.isLetterOrDigit()) {
            l++
            continue
        }
        if (!c2.isLetterOrDigit()) {
            r--
            continue
        }
        if (c1 != c2) {
            return false
        }
        l++
        r--
    }
    return true
}

fun main() {
    println("true == " + isPalindrome("A man, a plan, a canal: Panama"))
    println("false == " + isPalindrome("race a car"))
    println("false == " + isPalindrome("0P"))
}
