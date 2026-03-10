/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    count := 0                              // counter for smallest nodes in tree 
    stack := []*TreeNode{}                  // stack to push when travelling down/pop when counted
    curr := root                            

    for curr != nil || len(stack) != 0 {    // while travelling down or backtracking (in-order)
        for curr != nil {                   // if travelling down
            stack = append(stack, curr)     // push nodes to stack while travelling down
            curr = curr.Left
        }
        curr = stack[len(stack)-1]          // end of sub reached, pop last node from stack
        stack = stack[:len(stack)-1]
        count++                             // add to count for nth smallest node
        if count == k {                     // k smallest node reached, return that node's value
            return curr.Val
        }
        curr = curr.Right                   // process right subtree after left sub, root (in-order)
    }
    return 0                                // go required compile return (never reached)
}