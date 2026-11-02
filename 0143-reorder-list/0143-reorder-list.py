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
        slow, fast = head, head.next

        while fast and fast.next:
            fast = fast.next.next
            slow = slow.next

        curr = slow.next
        slow.next, prev = None, None

        while curr:
            nxt = curr.next
            curr.next = prev
            prev = curr
            curr = nxt

        a = head
        b = prev

        while b:
            a_nxt, b_nxt = a.next, b.next
            a.next = b
            b.next = a_nxt
            a, b = a_nxt, b_nxt





