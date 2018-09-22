#!/usr/bin/env python3

# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def removeElements(self, head, val):
        cur = head
        prev = None
        while cur != None:
            if cur.val != val:
                prev = cur
            else:
                if prev == None:
                    head = cur.next
                elif cur.next == None:
                    prev.next = None
                else:
                    prev.next = cur.next

            cur = cur.next
        return head
