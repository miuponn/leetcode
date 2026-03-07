/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    res := []int{}                                                  // res: slice of int
    if root == nil { return res }                                   // nil check

    q := []*TreeNode{root}                                          // init q
    for len(q) > 0 {                                                // while q not empty
        qLen := len(q)                                              // store var of current q length
        rMost := 0                                                  // var for rmost
        for i := 0; i < qLen; i++ {                                 // for each node in level
            node := q[0]                                            
            q = q[1:]
            if node != nil {                                        // if node not nil, process
                rMost = node.Val
                if node.Left != nil { q = append(q, node.Left) }
                if node.Right != nil { q = append(q, node.Right) }
            }
        }
        res = append(res, rMost)                                    // add last node from level as rMost
    }
    return res
}