package scala

/**
 * Definition for singly-linked list.
 */
class ListNode(_x: Int = 0, _next: ListNode = null) {
    var next: ListNode = _next
    var x: Int = _x
}

def rotateRight(head: ListNode, k: Int): ListNode = {
    if (head == null || head.next == null || k == 0) return head

    var length = 1
    var tail = head

    // Find the length of the list and the tail node
    while (tail.next != null) {
        tail = tail.next
        length += 1
    }

    val effectiveK = k % length
    if (effectiveK == 0) return head

    // Find the new tail node
    // The new tail is (length - effectiveK) nodes from the head
    // The new head is the next node of the new tail
    var newTail = head
    for (_ <- 1 until length - effectiveK) {
        newTail = newTail.next
    }
    val newHead = newTail.next

    //                newHead
    //                |
    // 1 -> 2 -> 3 -> 4 -> 5
    //           |
    //           newTail

    // Break the list at the new tail
    // and connect the tail to the head
    newTail.next = null
    tail.next = head

    newHead
}

def main(args: Array[String]): Unit = {
    val head = new ListNode(1, new ListNode(2, new ListNode(3, new ListNode(4, new ListNode(5)))))
    var rotatedList = rotateRight(head, 2)

    while (rotatedList != null) {
        print(rotatedList.x + " ")
        rotatedList = rotatedList.next
    }
}
