/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || root == p || root == q {            // Base case 1: end of tree + base case 2: found 1 target
        return root 
    }
    left := lowestCommonAncestor(root.Left, p, q)         // Recurse left and right subtrees
    right := lowestCommonAncestor(root.Right, p, q)
    if left != nil && right != nil {                      // Process left/right: non-nil, new root LCA
        return root
    }
    if left != nil { return left }                        // One nil, return LCA from that side's call
    return right
}