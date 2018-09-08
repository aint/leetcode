import scala.annotation.tailrec

object ReverseString {
    def reverseString(s: String): String = {
        val chars = s.toCharArray

        @tailrec
        def reverser(i: Int, j: Int): String = {
            if (i == s.length / 2) return chars.mkString
            val c = chars(i)
            chars(i) = chars(j)
            chars(j) = c

            reverser(i + 1, j - 1)
        }
        
        reverser(0, s.length - 1)
    }
}