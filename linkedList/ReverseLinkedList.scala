//Definition for singly-linked list.
class ListNode(var _x: Int = 0) {
   var next: ListNode = _
   var x: Int = _x
}

object Solution {
    def reverseList(head: ListNode): ListNode = {
        var tail: ListNode = null
        var cur = head

        while (cur != null) {
          val next = cur.next
          cur.next = tail
          tail = cur
          cur = next
        }
        
        tail
    }
}