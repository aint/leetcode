import scala.util.control.Breaks._
object Solution {

    def main(args: Array[String]): Unit = {
        val input = "abcabcbb"

        println("IN: " + input)

        println("OUT:" + lengthOfLongestSubstring(input))
      }

    def lengthOfLongestSubstring(s: String): Int = {
        var result = if (s.length() == 1) 1 else 0

        for (i <- 0 until s.length) {
            println(s" # i: $i (${s(i)})")
            breakable {

                var count = 1

                for (j <- i + 1 until s.length) {
                    println(s"        # j: $j (${s(j)})")
                    println("        " + s(i) + " == " + s(j))

                    if (s(j) == s(i)) {
                        println("        dup found")
                        println("        COUNT: " + count)
                        result = if (count > result) count else result
                        break
                    }

                    for (k <- i + 1 to j) {
                        println(s"            # k: $k (${s(k)})")
                        println("            " + s(j) + " == " + s(k))
                        if (s(j) == s(k) && j != k) {
                            println("            dup found")
                            println("            COUNT: " + count)
                            result = if (count > result) count else result
                            break
                        }
                    }

                    count += 1
                    result = if (count > result) count else result
                    println("        COUNT: " + count)

                }

            }
        }

        return result
    }
}