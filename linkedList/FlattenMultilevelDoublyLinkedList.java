// Definition for a Node.
class Node {
    public int val;
    public Node prev;
    public Node next;
    public Node child;

    public Node() {
    }

    public Node(int _val, Node _prev, Node _next, Node _child) {
        val = _val;
        prev = _prev;
        next = _next;
        child = _child;
    }
};

class Solution {
    public Node flatten(Node head) {
        rec(head);

        return null;
    }

    private void rec(Node node) {
        while (node != null) {
            System.out.println(node.val);

            if (node.child != null) {
                rec(node.child);
            }

            node = node.next;
        }
    }

}