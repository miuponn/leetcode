/**
 * Definition for singly-linked list.
 * class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */
public class Solution {
    public boolean hasCycle(ListNode head) {
        ListNode slow = head;                       // moves one step at a time
        ListNode fast = head;                       // moves two steps at a time

        if (head == null || head.next == null){     // empty or single-node list can't have a cycle
            return false;
        }

        while (fast != null && fast.next != null){  
            slow = slow.next;                       // move slow pointer by one, fast pointer by two
            fast = fast.next.next;

            if (slow == fast){                      // if both pointers eventually meet, a cycle exists
                return true;
            }
        }
        return false;                               
    }
}