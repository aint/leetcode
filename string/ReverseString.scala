import scala.annotation.tailrec

object Solution {
    def reverseString(s: String): String = {
        val chars = s.toList

        @tailrec
        def reverser(in: List[Char], out: List[Char]): String = {
            in match {
                case Nil => out.mkString
                case x :: tail => reverser(tail, x :: out)
            }
        }

        reverser(chars, Nil)        
    }
}