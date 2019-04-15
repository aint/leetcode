import scala.util.control.Breaks._
object Solution {

    def lengthOfLongestSubstring(s: String): Int = {
        var result = if (s.length() == 1) 1 else 0

        for (i <- 0 until s.length) {
            breakable {
                var count = 1

                for (j <- i + 1 until s.length) {
                    if (s(j) == s(i)) {
                        result = result.max(count)
                        break
                    }

                    for (k <- i + 1 to j) {
                        if (s(j) == s(k) && j != k) {
                            result = result.max(count)
                            break
                        }
                    }

                    count += 1
                    result = result.max(count)
                }
            }
        }

        return result
    }
}