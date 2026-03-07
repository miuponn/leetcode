/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    res := [][]int{}                               // init result slice
    if root == nil { return res }                  // nil check root

    q := []*TreeNode{root}                         // init queue, push root to q
    for len(q) > 0 {                               
        level := []int{}                           // init slice for level
        qLen := len(q)                             // var to store # of nodes to process next
        for i := 0; i < qLen; i++ {                // process all nodes for level
            node := q[0]                           // point to front, rest of q
            q = q[1:]                              
            level = append(level, node.Val)        // pop front, append to level slice
            
            if node.Left != nil {                  // push children to q for next level to process
                q = append(q, node.Left)          
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
        res = append(res, level)                   // add level slice to result slice
    }
    return res                                     // return result slice after all levels processed
}