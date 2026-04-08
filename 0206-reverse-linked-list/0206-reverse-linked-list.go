/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil {                        // edge case empty list return nil
        return nil
    }
    newhead := head                         // update new head with curr node
    if head.Next != nil {                   // if curr node not tail, recurse until tail found
        newhead = reverseList(head.Next)    // as call stack unwinds, tail node is passed down as new head
        head.Next.Next = head               // after recursive call, reverse pointer of curr node and next node
    }
    head.Next = nil                         // gets overwritten for all nodes except old head/new tail
    return newhead  
}