import scala.util.control.Breaks._
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
            breakable {

                var startIndex = 0
                var endIndex = 0

                for (j <- i + 1 until s.length) {
                    println(s"        $i,   $j")
                    println("        " + s(i) + " == " + s(j))

                    if (s(j) == s(i)) {
                        println("        dup found: " + s(j))
                        endIndex = j
                        println("        endIndex: " + j)

                        for (j <- i to 0 by -1) {
                            println(s"        $i,   $j")
                            println("        " + s(i) + " == " + s(j))

                            if (s(j) == s(i)) {
                                println("        dup found: " + s(j))
                                startIndex = j
                                println("        startIndex: " + j)

                                var indexDiff = if (endIndex - 1 == startIndex) 0 else endIndex - startIndex
                                println("        indexDiff: " + indexDiff)

                                result = if (indexDiff > result) indexDiff else result

                                break
                            }
                        }
                    }
                }

            }
        }

        return result
    }
}