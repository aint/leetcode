object Solution {

    def main(args: Array[String]): Unit = {
        val input = "abcabcbb"

        println("IN: " + input)

        println("OUT:" + lengthOfLongestSubstring(input))
      }

    def lengthOfLongestSubstring(s: String): Int = {
        var start = 0
        var result = 0

        for (i <- 0 until s.length) {
            println(" # " + i + " is "+ s(i))
            // check uniquest in some seq from `start` to `i`
            for (j <- start until i) {
                if (s(j) == s(i)) {
                    println("        dup found: " + s(j))
                    result = if (result > i - start) result else i - start
                    start = i
                }
            }
        }

        return result
    }
}