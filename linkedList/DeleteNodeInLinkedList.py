# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def deleteNode(self, node):
        if node.next is not None:
            node.val = node.next.val
            node.next = node.next.next
