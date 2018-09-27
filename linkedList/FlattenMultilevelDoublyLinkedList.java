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
        Node prev = head;
        while (head != null) {
            System.out.println(head.val);

            Node child = head.child;
            head.child = null;
            if (child != null) {
                System.out.println("s " + head.val);
                Node tmpNext = head.next;
                head.next = child;
                child.prev = head;

                while(child.next != null) {
                    child = child.next;
                }

                System.out.println("e " + child.val);
                child.next = tmpNext;
                if (tmpNext != null) {
                    tmpNext.prev = child;
                }
            }

            head = head.next;
        }
        return prev;
    }

}