# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reorderList(self, head: ListNode | None) -> None:
        """
        Do not return anything, modify head in-place instead.
        """
        slow, fast = head, head.next        # declare fast, slow pointers to find center of list

        while fast and fast.next:           # while fast node is not last node in list:
            fast = fast.next.next               # move fast pointer forward twice, slow pointer once
            slow = slow.next                       

        curr = slow.next                    # declare first node of second half
        slow.next, prev = None, None        # split list into first and second half, initate prev pointer

        while curr:                         # while node of second half is not null:
            nxt = curr.next                     # declare temp pointer for next node
            curr.next = prev                    # point curr node at prev node (reversing pointer)
            prev = curr                         # move prev pointer forward
            curr = nxt                          # move curr pointer forward

        a, b = head, prev                   # declare pointers to first nodes of first half, second half

        while b:                            # while second half node is not null:
            a_nxt, b_nxt = a.next, b.next       # declare temp pointers for next nodes
            a.next = b                          # point first half node to second half node
            b.next = a_nxt                      # point second half node to next first half node
            a, b = a_nxt, b_nxt                 # move pointers forward





