/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode dummy = new ListNode(0);                   // dummy node to simplify list construction
        ListNode curr = dummy;                              // current pointer to build the result list
        int carry = 0;                                      // carry value for sums greater than 9

        while (l1 != null || l2 != null || carry > 0){      // continue while there are nodes to process or carry is non-zero
            if (l1 != null) {                               // if l1 node exists, add its value to carry
                carry += l1.val;                                
                l1 = l1.next;                               // move l1 pointer to next node
            }
            if (l2 != null) {                               // if l2 node exists, add its value to carry
                carry += l2.val;
                l2 = l2.next;                               // move l2 pointer to next node
            }   
            curr.next = new ListNode(carry % 10);           // create new node with the last digit of carry
            curr = curr.next;                               // move current pointer forward
            carry /= 10;                                    // update carry for next iteration
        }
        return dummy.next;                                  // return head of the resulting sum list
    }
}